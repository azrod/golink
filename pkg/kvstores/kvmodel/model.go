package kvmodel

import (
	"context"

	"github.com/azrod/golink/models"
)

type KV interface {
	Close() error

	List(ctx context.Context, prefix string) (keys []string, err error)
	Get(ctx context.Context, key string, value models.Model) error
	Set(ctx context.Context, key string, value models.Model) error
	Delete(ctx context.Context, key string) error
}
