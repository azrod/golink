package clients

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"os"

	"github.com/sethvargo/go-envconfig"

	"github.com/azrod/golink/pkg/clients/clientmodel"
	"github.com/azrod/golink/pkg/clients/redis"
)

type (
	Settings struct {
		Redis struct {
			// DBRedisAddress is the address of the redis database
			Address string `yaml:"dbRedisAddress" env:"DB_REDIS_ADDRESS,default=localhost:6379"`

			// Password is the password of the redis database
			Password string `yaml:"dbRedisPassword" env:"DB_REDIS_PASSWORD"`

			// Username is the username of the redis database
			Username string `yaml:"dbRedisUsername" env:"DB_REDIS_USERNAME"`

			// DB is the database number of the redis database
			DB int `yaml:"dbRedisDb" env:"DB_REDIS_DB,default=0"`

			// MaxRetries is the maximum number of retries before giving up
			MaxRetries int `yaml:"dbRedisMaxRetries" env:"DB_REDIS_MAX_RETRIES,default=3"`

			// DialTimeout is the maximum number of retries before giving up
			DialTimeout int `yaml:"dbDialTimeout" env:"DB_REDIS_DIAL_TIMEOUT,default=5"`

			// ReadTimeout is the maximum number of retries before giving up
			ReadTimeout int `yaml:"dbReadTimeout" env:"DB_REDIS_READ_TIMEOUT,default=3"`

			CertFile string `yaml:"dbRedisCertFile" env:"DB_REDIS_CERT_FILE"`
			KeyFile  string `yaml:"dbRedisKeyFile" env:"DB_REDIS_KEY_FILE"`
			CAFile   string `yaml:"dbRedisCAFile" env:"DB_REDIS_CA_FILE"`

			// TODO add tls config
		} `yaml:"redis"`
	}
)

// NewClient creates a new client.
func NewClient(ctx context.Context, cfg Settings) (clientmodel.ClientDB, error) {
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	switch {
	case cfg.Redis.Address != "":
		if cfg.Redis.CertFile != "" && cfg.Redis.KeyFile != "" && cfg.Redis.CAFile != "" {
			// open cert files
			serverTLSCert, err := tls.LoadX509KeyPair(cfg.Redis.CertFile, cfg.Redis.KeyFile)
			if err != nil {
				log.Fatalf("Error loading certificate and key file: %v", err)
			}

			// open ca file
			caCert, err := os.ReadFile(cfg.Redis.CAFile)
			if err != nil {
				log.Fatalf("Error reading ca file: %v", err)
			}

			// create cert pool
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)

			// create tls config
			tlsConfig := &tls.Config{ //nolint:gosec
				Certificates: []tls.Certificate{serverTLSCert},
				ClientCAs:    caCertPool,
			}

			return redis.New(redis.Settings{
				Address:     cfg.Redis.Address,
				Username:    cfg.Redis.Username,
				Password:    cfg.Redis.Password,
				DB:          cfg.Redis.DB,
				MaxRetries:  cfg.Redis.MaxRetries,
				DialTimeout: cfg.Redis.DialTimeout,
				ReadTimeout: cfg.Redis.ReadTimeout,
				TLSConfig:   tlsConfig,
			})
		}

		return redis.New(redis.Settings{
			Address:     cfg.Redis.Address,
			Username:    cfg.Redis.Username,
			Password:    cfg.Redis.Password,
			DB:          cfg.Redis.DB,
			MaxRetries:  cfg.Redis.MaxRetries,
			DialTimeout: cfg.Redis.DialTimeout,
			ReadTimeout: cfg.Redis.ReadTimeout,
		})

	default:
		return nil, errors.New("no client found")
	}
}
