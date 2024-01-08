package sb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/kvstores/kvmodel"
)

// GetLinkByID gets a link from the database.
func (c Client) GetLinkByID(ctx context.Context, linkID models.LinkID, namespace string) (link models.Link, err error) {
	key := linkID.String()

	if !strings.Contains(linkID.String(), linkKey) {
		key = fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String())
	}

	if err := c.c.Get(ctx, key, &link); err != nil {
		if errors.Is(err, kvmodel.ErrNotFound) {
			return models.Link{}, fmt.Errorf("link %s %w", linkID.String(), models.ErrNotFound)
		}
		return models.Link{}, err
	}

	return link, nil
}

// GetLinkByPath gets a link from the database.
func (c Client) GetLinkByPath(ctx context.Context, pathn, namespace string) (models.Link, error) {
	keys, err := c.c.List(ctx, fmt.Sprintf("%s%s", linkKey, namespace))
	if err != nil {
		return models.Link{}, err
	}

	for _, key := range keys {
		link, err := c.GetLinkByID(ctx, models.LinkID(key), namespace)
		if err != nil {
			continue
		}

		if link.SourcePath == pathn {
			return link, nil
		}
	}

	return models.Link{}, fmt.Errorf("link path %s is %w", pathn, models.ErrNotFound)
}

// GetLinkByName gets a link from the database.
func (c Client) GetLinkByName(ctx context.Context, name, namespace string) (models.Link, error) {
	links, err := c.ListLinks(ctx, namespace)
	if err != nil {
		return models.Link{}, err
	}

	for _, link := range links {
		if link.Name == name {
			return link, nil
		}
	}

	return models.Link{}, fmt.Errorf("link name %s is %w", name, models.ErrNotFound)
}

// UpdateLink updates a link in the database.
func (c Client) UpdateLink(ctx context.Context, link models.LinkRequest, linkID models.LinkID, namespace string) (models.Link, error) {
	l, err := c.GetLinkByID(ctx, linkID, namespace)
	if err != nil {
		return models.Link{}, err
	}

	l.LinkRequest = link

	if err := l.Validate(); err != nil {
		return models.Link{}, err
	}

	return l, c.c.Set(ctx, fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String()), &l)
}

// DeleteLink deletes a link from the database.
func (c Client) DeleteLink(ctx context.Context, linkID models.LinkID, namespace string) error {
	return c.c.Delete(ctx, fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String()))
}

// CreateLink creates a link in the database.
func (c Client) CreateLink(ctx context.Context, link models.LinkRequest, namespace string) (models.Link, error) {
	l := models.Link{
		LinkRequest: link,
		NameSpace:   namespace,
	}

	if err := l.Validate(); err != nil {
		return models.Link{}, err
	}

	return l, c.c.Set(ctx, fmt.Sprintf("%s%s:%s", linkKey, namespace, l.ID.String()), &l)
}

// CountLinks counts all links for a namespace in the database.
func (c Client) CountLinks(ctx context.Context, namespace string) (int, error) {
	keys, err := c.c.List(ctx, fmt.Sprintf("%s%s", linkKey, namespace))
	if err != nil {
		return 0, err
	}

	return len(keys), nil
}

// ListLinks lists all links in the database.
func (c Client) ListLinks(ctx context.Context, namespace string) ([]models.Link, error) {
	links := make([]models.Link, 0)
	keys, err := c.c.List(ctx, fmt.Sprintf("%s%s", linkKey, namespace))
	if err != nil {
		return []models.Link{}, err
	}

	for _, key := range keys {
		link, err := c.GetLinkByID(ctx, models.LinkID(key), namespace)
		if err != nil {
			continue
		}

		links = append(links, link)
	}

	return links, nil
}
