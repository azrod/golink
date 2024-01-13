package sb

import (
	"fmt"
	"log"

	"github.com/azrod/golink/pkg/kvstores/kvmodel"
	"github.com/azrod/golink/pkg/kvstores/local"
	"github.com/azrod/golink/pkg/kvstores/redis"
)

const (
	// Struct Path KEY.
	linkKey      = "link:"
	labelKey     = "label:"
	NamespaceKey = "ns:"
)

const (
	TypeOfBackendRedis TypeOfBackend = "redis"
	TypeOfBackendLocal TypeOfBackend = "local"
)

type (
	TypeOfBackend string
	Config        struct {
		Type  TypeOfBackend `default:"local"`
		Redis redis.Option
		Local local.Option
	}
	Client struct {
		c kvmodel.KV
	}
)

func New(config Config) (client Client, err error) {
	switch config.Type {
	case TypeOfBackendRedis:
		c, err := redis.New(config.Redis)
		if err != nil {
			return Client{}, err
		}

		client = Client{c}
	case TypeOfBackendLocal:
		c, err := local.New(config.Local)
		if err != nil {
			return Client{}, err
		}

		client = Client{c}
	default:
		return Client{}, fmt.Errorf("storage backend %s is not supported", config.Type)
	}

	log.Default().Printf("Successfully connected to the storage backend: %s", config.Type)
	return client, err
}

func (c Client) Close() error {
	return c.c.Close()
}
