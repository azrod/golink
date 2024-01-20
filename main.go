package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/azrod/golink/api"
	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/config"
	"github.com/azrod/golink/pkg/sb"
	"github.com/azrod/golink/short"

	_ "github.com/azrod/golink/docs/echosimple"
)

var (
	globalEchoLogLevel = log.DEBUG
	version            = "dev"
	ready              = false
)

func main() {
	ctx := context.Background()

	// * Load config
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	// * Setup Storage Backend
	db, err := sb.New(cfg.Storage)
	if err != nil {
		panic(err)
	}

	// * Check if Namespace default exists
	if _, err := db.GetNamespace(ctx, "default"); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			if _, err := db.CreateNamespace(ctx, models.NamespaceRequest{
				Name: "default",
			}); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	type serverCfg struct {
		Address string
		Port    int
	}

	httpServers := make(map[serverCfg]*echo.Echo)

	// * Health server
	healthServer := echo.New()
	healthServer.HideBanner = true
	healthServer.Logger.SetLevel(globalEchoLogLevel)
	healthServer.Use(middleware.Recover())
	healthServer.GET("/health", func(c echo.Context) error {
		// TODO add status check for DB
		if !ready {
			return c.String(http.StatusServiceUnavailable, "Not ready")
		}
		return c.String(http.StatusOK, "OK")
	})
	httpServers[serverCfg{
		Address: cfg.Health.Address,
		Port:    cfg.Health.Port,
	}] = healthServer

	// * App server
	appServer := echo.New()
	appServer.HideBanner = true
	appServer.Logger.SetLevel(globalEchoLogLevel)
	appServer.Use(middleware.Logger())
	appServer.Use(middleware.Recover())
	appServer.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("version", version)
			return next(c)
		}
	}))
	ui := appServer.Group("/u")
	ui.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	ui.HEAD("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	if version != "dev" {
		ui.Static("/*", "www")
	} else {
		ui.Static("/*", "ui/dist/")
	}

	api.NewHandlers(db, appServer)
	short.NewHandlers(db, appServer)
	httpServers[serverCfg{
		Address: cfg.App.Address,
		Port:    cfg.App.Port,
	}] = appServer

	// * Start HTTP servers
	for cfg, server := range httpServers {
		go func(server *echo.Echo, cfg serverCfg) {
			server.Logger.Infof("Starting server on %s:%d", cfg.Address, cfg.Port)
			if err := server.Start(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
				server.Logger.Fatal("shutting down the server")
			}
		}(server, cfg)
	}

	time.Sleep(3 * time.Second)
	ready = true

	// Wait for interrupt signal (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// ! Graceful shutdown
	for _, server := range httpServers {
		server.Logger.Info("Shutting down server")
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctxTimeout); err != nil {
			server.Logger.Error(err)
		}
		cancel()
	}

	// * Close Storage Backend
	if err := db.Close(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
