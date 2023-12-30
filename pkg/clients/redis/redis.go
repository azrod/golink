package redis

import (
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/azrod/golink/pkg/clients/clientmodel"
)

// _ clients.Client   = (*client)(nil).
var _ clientmodel.ClientDB = (*Client)(nil)

const (
	// Struct Path KEY redis.
	linkKey      = "link:"
	labelKey     = "label:"
	NamespaceKey = "ns:"
)

type (
	Settings struct {
		Address      string `yaml:"address"`      // host:port address.
		Username     string `yaml:"username"`     // Optional username.
		Password     string `yaml:"password"`     // Optional password.
		MaxRetries   int    `yaml:"maxRetries"`   // Maximum number of retries before giving up.
		DialTimeout  int    `yaml:"dialTimeout"`  // in seconds
		ReadTimeout  int    `yaml:"readTimeout"`  // in seconds
		WriteTimeout int    `yaml:"writeTimeout"` // in seconds
		DB           int    `yaml:"db"`           // Database to be selected after connecting to the server.
		TLSConfig    *tls.Config
	}
	Client struct {
		c *redis.Client
	}
)

func New(s Settings) (Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:                  s.Address,
		Username:              s.Username,
		Password:              s.Password,
		MaxRetries:            s.MaxRetries,
		DB:                    s.DB,
		ContextTimeoutEnabled: true,
		DialTimeout:           time.Duration(s.DialTimeout) * time.Second,
		ReadTimeout:           time.Duration(s.ReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(s.WriteTimeout) * time.Second,
		TLSConfig:             s.TLSConfig,
	})

	return Client{client}, nil
}
