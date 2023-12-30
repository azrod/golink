package redis

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/azrod/golink/models"
)

// GetLinkByID gets a link from the database.
func (c Client) GetLinkByID(ctx context.Context, linkID models.LinkID, namespace string) (models.Link, error) {
	if !strings.Contains(linkID.String(), linkKey) {
		linkID = models.LinkID(fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String()))
	}

	v, err := c.c.Get(ctx, linkID.String()).Result()
	if err != nil {
		return models.Link{}, err
	}

	// v is JSON string
	var link models.Link
	if err := link.UnmarshalJSON([]byte(v)); err != nil {
		return models.Link{}, err
	}

	return link, nil
}

// GetLinkByPath gets a link from the database.
func (c Client) GetLinkByPath(ctx context.Context, path, namespace string) (models.Link, error) {
	iter := c.c.Scan(ctx, 0, fmt.Sprintf("%s%s:*", linkKey, namespace), 0).Iterator()
	for iter.Next(ctx) {
		link, err := c.GetLinkByID(ctx, models.LinkID(iter.Val()), namespace)
		if err != nil {
			continue
		}

		if link.SourcePath == path {
			return link, nil
		}
	}

	return models.Link{}, fmt.Errorf("link path %s is %w", path, models.ErrNotFound)
}

// GetLinkByName gets a link from the database.
func (c Client) GetLinkByName(ctx context.Context, name, namespace string) (models.Link, error) {
	iter := c.c.Scan(ctx, 0, fmt.Sprintf("%s%s:*", linkKey, namespace), 0).Iterator()
	for iter.Next(ctx) {
		link, err := c.GetLinkByID(ctx, models.LinkID(iter.Val()), namespace)
		if err != nil {
			continue
		}

		log.Default().Printf("link.Name: %s, name: %s", link.Name, name)

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

	b, err := l.MarshalJSON()
	if err != nil {
		return models.Link{}, err
	}

	return l, c.c.Set(ctx, fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String()), string(b), 0).Err()
}

// DeleteLink deletes a link from the database.
func (c Client) DeleteLink(ctx context.Context, linkID models.LinkID, namespace string) error {
	return c.c.Del(ctx, fmt.Sprintf("%s%s:%s", linkKey, namespace, linkID.String())).Err()
}

// CreateLink creates a link in the database.
func (c Client) CreateLink(ctx context.Context, link models.LinkRequest, namespace string) (models.Link, error) {
	l := models.Link{
		LinkRequest: link,
		NameSpace:   namespace,
	}

	b, err := l.MarshalJSON()
	if err != nil {
		return models.Link{}, err
	}
	return l, c.c.Set(ctx, fmt.Sprintf("%s%s:%s", linkKey, l.NameSpace, l.ID.String()), string(b), 0).Err()
}

// ListLinks lists all links in the database.
func (c Client) ListLinks(ctx context.Context, namespace string) ([]models.Link, error) {
	links := make([]models.Link, 0)
	iter := c.c.Scan(ctx, 0, fmt.Sprintf("%s%s:*", linkKey, namespace), 0).Iterator()
	for iter.Next(ctx) {
		link, err := c.GetLinkByID(ctx, models.LinkID(iter.Val()), namespace)
		if err != nil {
			continue
		}

		links = append(links, link)
	}

	return links, nil
}
