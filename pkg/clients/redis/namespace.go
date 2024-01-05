package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"

	"github.com/azrod/golink/models"
)

// GetNamespace gets a Namespace from the database.
func (c Client) GetNamespace(ctx context.Context, namespace string) (models.Namespace, error) {
	nsKey := namespace
	if !strings.Contains(nsKey, NamespaceKey) {
		nsKey = NamespaceKey + nsKey
	}

	v, err := c.c.Get(ctx, nsKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return models.Namespace{}, fmt.Errorf("namespace %s %w", namespace, models.ErrNotFound)
		}
		return models.Namespace{}, err
	}

	// v is JSON string
	var Namespace models.Namespace
	if err := Namespace.UnmarshalJSON([]byte(v)); err != nil {
		return models.Namespace{}, err
	}

	links, err := c.ListLinks(ctx, Namespace.Name)
	if err != nil {
		return models.Namespace{}, err
	}

	Namespace.Links = links

	return Namespace, nil
}

// DeleteNamespace deletes a Namespace from the database.
func (c Client) DeleteNamespace(ctx context.Context, namespace string) error {
	ns := namespace

	if !strings.Contains(namespace, NamespaceKey) {
		namespace = NamespaceKey + namespace
	}

	// Check if namespace exists
	if _, err := c.GetNamespace(ctx, ns); err != nil {
		return err
	}

	// Check if namespace is empty
	if links, err := c.ListLinks(ctx, ns); err != nil || len(links) > 0 {
		if err != nil {
			return err
		}
		// Found links - return error
		return fmt.Errorf("namespace %s is not empty\nFound %d link(s)", ns, len(links))
	}

	return c.c.Del(ctx, namespace).Err()
}

// CreateNamespace creates a Namespace in the database.
func (c Client) CreateNamespace(ctx context.Context, namespace models.NamespaceRequest) (models.Namespace, error) {
	ns := models.Namespace{
		NamespaceRequest: namespace,
		Enabled:          models.Enabled(true),
	}

	b, err := ns.MarshalJSON()
	if err != nil {
		return models.Namespace{}, err
	}

	// Check if namespace already exists
	if _, err := c.GetNamespace(ctx, namespace.Name); err == nil {
		return models.Namespace{}, fmt.Errorf("namespace %s already exists", namespace.Name)
	}

	if err := c.c.Set(ctx, NamespaceKey+namespace.Name, string(b), 0).Err(); err != nil {
		return models.Namespace{}, err
	}

	return c.GetNamespace(ctx, namespace.Name)
}

// ListNamespaces lists all Namespaces in the database.
func (c Client) ListNamespaces(ctx context.Context) ([]models.Namespace, error) {
	var Namespaces []models.Namespace
	iter := c.c.Scan(ctx, 0, NamespaceKey+"*", 0).Iterator()
	for iter.Next(ctx) {
		Namespace, err := c.GetNamespace(ctx, iter.Val())
		if err != nil {
			continue
		}

		Namespaces = append(Namespaces, Namespace)
	}

	for i, ns := range Namespaces {
		links, err := c.ListLinks(ctx, ns.Name)
		if err != nil {
			continue
		}

		Namespaces[i].Links = links
	}

	return Namespaces, nil
}

// // AddLinkToNamespace adds a link to a Namespace in the database.
// func (c Client) AddLinkToNamespace(ctx context.Context, namespace string, link models.LinkRequest) error {
// 	_, err := c.GetNamespace(ctx, namespace)
// 	if err != nil {
// 		return err
// 	}

// 	if namespace == "" {
// 		return fmt.Errorf("%s %w", namespace, models.ErrIsEmpty)
// 	}

// 	return c.CreateLink(ctx, link, namespace)
// }

// // RemoveLinkFromNamespace removes a link from a Namespace in the database.
// func (c Client) RemoveLinkFromNamespace(ctx context.Context, namespace string, linkID models.LinkID) error {
// 	_, err := c.GetNamespace(ctx, namespace)
// 	if err != nil {
// 		return err
// 	}

// 	return c.DeleteLink(ctx, linkID, namespace)
// }
