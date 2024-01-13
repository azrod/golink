package sb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/kvstores/kvmodel"
)

// GetNamespace gets a namespace from the database.
func (c Client) GetNamespace(ctx context.Context, namespace string) (ns models.Namespace, err error) {
	key := namespace

	if !strings.Contains(namespace, NamespaceKey) {
		key = fmt.Sprintf("%s%s", NamespaceKey, namespace)
	}

	if err := c.c.Get(ctx, key, &ns); err != nil {
		if errors.Is(err, kvmodel.ErrNotFound) {
			return models.Namespace{}, fmt.Errorf("namespace %s is %w", namespace, models.ErrNotFound)
		}
		return models.Namespace{}, err
	}

	links, err := c.ListLinks(ctx, ns.Name)
	if err != nil {
		return models.Namespace{}, err
	}

	ns.Links = links

	return ns, nil
}

// CreateNamespace creates a namespace in the database.
func (c Client) CreateNamespace(ctx context.Context, namespace models.NamespaceRequest) (ns models.Namespace, err error) {
	ns = models.Namespace{
		NamespaceRequest: namespace,
		Enabled:          models.Enabled(true),
	}

	if err := ns.Validate(); err != nil {
		return models.Namespace{}, err
	}

	// Check if namespace exists
	if _, err := c.GetNamespace(ctx, ns.Name); err == nil {
		return models.Namespace{}, fmt.Errorf("namespace %s %w", ns.Name, models.ErrAlreadyExists)
	}

	if err := c.c.Set(ctx, NamespaceKey+ns.Name, &ns); err != nil {
		return models.Namespace{}, err
	}

	return c.GetNamespace(ctx, ns.Name)
}

// DeleteNamespace deletes a namespace from the database.
func (c Client) DeleteNamespace(ctx context.Context, namespace string) error {
	// Check if namespace exists
	if _, err := c.GetNamespace(ctx, NamespaceKey+namespace); err != nil {
		return err
	}

	// Check if namespace is empty
	if links, err := c.CountLinks(ctx, namespace); err != nil || links > 0 {
		if err != nil {
			return err
		}
		// Found links - return error
		return fmt.Errorf("namespace %s is not empty\nFound %d link(s)", namespace, links)
	}

	return c.c.Delete(ctx, NamespaceKey+namespace)
}

// ListNamespaces lists all namespaces in the database.
func (c Client) ListNamespaces(ctx context.Context) (nss []models.Namespace, err error) {
	namespaces, err := c.c.List(ctx, NamespaceKey)
	if err != nil {
		return nil, err
	}

	// add mutex write for nss slice
	mutex := sync.Mutex{}

	// new wait group
	wg := sync.WaitGroup{}
	chErr := make(chan error)

	nss = make([]models.Namespace, len(namespaces))

	for i, namespace := range namespaces {
		wg.Add(1)
		go func(namespace string, index int) {
			defer wg.Done()

			ns, err := c.GetNamespace(ctx, namespace)
			if err != nil {
				chErr <- err
				return
			}

			mutex.Lock()
			nss[index] = ns
			mutex.Unlock()
		}(namespace, i)
	}

	go func() {
		for err := range chErr {
			log.Default().Printf("Failed to get namespace: %s", err)
			return
		}
	}()

	// wait for all goroutines to finish
	wg.Wait()
	close(chErr)

	return nss, nil
}
