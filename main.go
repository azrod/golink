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
	"github.com/sethvargo/go-envconfig"

	"github.com/azrod/golink/api"
	"github.com/azrod/golink/pkg/clients"
	"github.com/azrod/golink/short"

	_ "github.com/azrod/golink/docs/echosimple"
)

var (
	globalEchoLogLevel = log.DEBUG

	ready = false
)

type config struct {
	ServerPort int    `env:"SERVER_PORT,default=8081"`
	ServerHost string `env:"SERVER_HOST,default=localhost"`
	ServerURL  string `env:"SERVER_URL,default=http://localhost:8081"`

	HealthServerPort int    `env:"HEALTH_SERVER_PORT,default=8082"`
	HealthServerHost string `env:"HEALTH_SERVER_HOST,default=localhost"`
}

func main() {
	ctx := context.Background()

	cfg := config{}

	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(err)
	}

	db, err := clients.NewClient(ctx, clients.Settings{})
	if err != nil {
		panic(err)
	}

	type serverCfg struct {
		Host string
		Port int
	}

	httpServers := make(map[serverCfg]*echo.Echo)

	// * Health server
	healthServer := echo.New()
	healthServer.HideBanner = true
	healthServer.Logger.SetLevel(globalEchoLogLevel)
	healthServer.Use(middleware.Recover())
	healthServer.GET("/health", func(c echo.Context) error {
		if !ready {
			return c.String(http.StatusServiceUnavailable, "Not ready")
		}
		return c.String(http.StatusOK, "OK")
	})
	httpServers[serverCfg{
		Host: cfg.HealthServerHost,
		Port: cfg.HealthServerPort,
	}] = healthServer

	// * App server
	appServer := echo.New()
	appServer.HideBanner = true
	appServer.Logger.SetLevel(globalEchoLogLevel)
	appServer.Use(middleware.Logger())
	appServer.Use(middleware.Recover())

	ui := appServer.Group("/u")
	ui.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	ui.HEAD("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	ui.Static("/*", "ui/dist")

	api.NewHandlers(db, appServer)
	short.NewHandlers(db, appServer)
	httpServers[serverCfg{
		Host: cfg.ServerHost,
		Port: cfg.ServerPort,
	}] = appServer

	// * Start HTTP servers
	for cfg, server := range httpServers {
		go func(server *echo.Echo, cfg serverCfg) {
			server.Logger.Infof("Starting server on %s:%d", cfg.Host, cfg.Port)
			if err := server.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	os.Exit(0)
}
