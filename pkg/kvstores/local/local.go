package local

import (
	"context"
	"fmt"
	"strings"

	"go.etcd.io/bbolt"

	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/kvstores/kvmodel"
)

var _ kvmodel.KV = (*Client)(nil)

type (
	Client struct {
		c *bbolt.DB
	}

	Option struct {
		Path string `default:"./"`
	}
)

var dbNames = []string{
	"link",
	"label",
	"namespace",
}

const (
	// Struct Path KEY.
	linkKey      = "link:"
	labelKey     = "label:"
	NamespaceKey = "ns:"
)

func New(opt Option) (Client, error) {
	if opt.Path[len(opt.Path)-1] != '/' {
		opt.Path += "/"
	}

	db, err := bbolt.Open(fmt.Sprintf("%sgolink.db", opt.Path), 0o600, nil)
	if err != nil {
		return Client{}, fmt.Errorf("could not open local db: %w", err)
	}

	if err := db.Update(func(tx *bbolt.Tx) error {
		for _, name := range dbNames {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return fmt.Errorf("could not create %s local db: %w", name, err)
			}
		}
		return nil
	}); err != nil {
		return Client{}, fmt.Errorf("could not create local db: %w", err)
	}

	return Client{db}, nil
}

func (c Client) Close() error {
	return c.c.Close()
}

func getBucket(key string) string {
	switch {
	case strings.Contains(key, linkKey):
		return "link"
	case strings.Contains(key, labelKey):
		return "label"
	case strings.Contains(key, NamespaceKey):
		return "namespace"
	default:
		return ""
	}
}

// Get gets a value from the database.
func (c Client) Get(_ context.Context, key string, value models.Model) error {
	return c.c.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(getBucket(key)))
		if b == nil {
			return kvmodel.ErrNotFound
		}

		v := b.Get([]byte(key))
		if v == nil {
			return kvmodel.ErrNotFound
		}

		return value.UnmarshalJSON(v)
	})
}

// Set sets a value in the database.
func (c Client) Set(_ context.Context, key string, value models.Model) error {
	return c.c.Update(func(tx *bbolt.Tx) error {
		// Check if key is not empty
		if key == "" {
			return kvmodel.ErrEmptyKey
		}

		b := tx.Bucket([]byte(getBucket(key)))
		if b == nil {
			return kvmodel.ErrNotFound
		}

		vJSON, err := value.MarshalJSON()
		if err != nil {
			return err
		}

		return b.Put([]byte(key), vJSON)
	})
}

// Delete deletes a value from the database.
func (c Client) Delete(_ context.Context, key string) error {
	// Check if key is not empty
	if key == "" {
		return kvmodel.ErrEmptyKey
	}
	return c.c.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(getBucket(key)))
		if b == nil {
			return kvmodel.ErrNotFound
		}

		return b.Delete([]byte(key))
	})
}

// List lists all the keys from the database.
func (c Client) List(_ context.Context, prefix string) (keys []string, err error) {
	return keys, c.c.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(getBucket(prefix)))
		if b == nil {
			return kvmodel.ErrNotFound
		}

		c := b.Cursor()

		for k, _ := c.Seek([]byte(prefix)); k != nil && strings.HasPrefix(string(k), prefix); k, _ = c.Next() {
			keys = append(keys, string(k))
		}

		return nil
	})
}
