package redis

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/kvstores/kvmodel"
)

var _ kvmodel.KV = (*Client)(nil)

type (
	Client struct {
		c *redis.Client
	}
	Option struct {
		// DBRedisAddress is the address of the redis database
		Address string `default:"localhost:6379"`

		// Password is the password of the redis database
		Password string

		// Username is the username of the redis database
		Username string

		// DB is the database number of the redis database
		DB int `default:"0"`

		// MaxRetries is the maximum number of retries before giving up
		MaxRetries int `default:"3"`

		// DialTimeout is the maximum number of retries before giving up
		DialTimeout int `default:"5"`

		// ReadTimeout is the maximum number of retries before giving up
		ReadTimeout int `default:"3"`

		// WriteTimeout is the maximum number of retries before giving up
		WriteTimeout int `default:"3"`

		CertFile string
		KeyFile  string
		CAFile   string
	}
)

func New(opt Option) (Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:                  opt.Address,
		Username:              opt.Username,
		Password:              opt.Password,
		MaxRetries:            opt.MaxRetries,
		DB:                    opt.DB,
		ContextTimeoutEnabled: true,
		DialTimeout:           time.Duration(opt.DialTimeout) * time.Second,
		ReadTimeout:           time.Duration(opt.ReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(opt.WriteTimeout) * time.Second,
		// TLSConfig:             opt.TLSConfig,
	})

	return Client{client}, nil
}

func (c Client) Set(ctx context.Context, key string, value models.Model) error {
	// Check if key is not empty
	if key == "" {
		return kvmodel.ErrEmptyKey
	}

	vJSON, err := value.MarshalJSON()
	if err != nil {
		return err
	}

	return c.c.Set(ctx, key, string(vJSON), 0).Err()
}

func (c Client) Get(ctx context.Context, key string, value models.Model) error {
	// Check if key is not empty
	if key == "" {
		return kvmodel.ErrEmptyKey
	}

	v, err := c.c.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return kvmodel.ErrNotFound
		}
		return err
	}

	return value.UnmarshalJSON([]byte(v))
}

// List lists all the keys from the database.
func (c Client) List(ctx context.Context, prefix string) (keys []string, err error) {
	// Check if prefix is not empty
	if prefix == "" {
		return nil, kvmodel.ErrEmptyPrefix
	}

	pattern := func() string {
		if strings.HasSuffix(prefix, ":") {
			return prefix + "*"
		}
		return prefix + ":*"
	}()

	return c.c.Keys(ctx, pattern).Result()
}

// Delete deletes the value for the given key.
func (c Client) Delete(ctx context.Context, key string) error {
	// Check if key is not empty
	if key == "" {
		return kvmodel.ErrEmptyKey
	}

	return c.c.Del(ctx, key).Err()
}

// Close closes the database connection.
func (c Client) Close() error {
	return c.c.Close()
}
