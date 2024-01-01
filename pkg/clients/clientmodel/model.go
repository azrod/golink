package clientmodel

import (
	"context"

	"github.com/azrod/golink/models"
)

type (
	Client interface {
		// New creates a new client
		New() (ClientDB, error)
	}

	ClientDB interface {
		// TODO add Close method

		// * Link
		// GetLinkByID gets a link from the database
		GetLinkByID(ctx context.Context, linkID models.LinkID, namespace string) (models.Link, error)

		// GetLinkByPath gets a link from the database
		GetLinkByPath(ctx context.Context, pathn, namespace string) (models.Link, error)

		// GetLinkByName gets a link from the database
		GetLinkByName(ctx context.Context, name, namespace string) (models.Link, error)

		// UpdateLink updates a link in the database
		UpdateLink(ctx context.Context, link models.LinkRequest, linkID models.LinkID, namespace string) (models.Link, error)

		// DeleteLink deletes a link from the database
		DeleteLink(ctx context.Context, linkID models.LinkID, namespace string) error

		// CreateLink creates a link in the database
		CreateLink(ctx context.Context, link models.LinkRequest, namespace string) (models.Link, error)

		// ListLinks lists all links in the database
		ListLinks(ctx context.Context, namespace string) ([]models.Link, error)

		// * Label
		// GetLabelByID gets a label from the database
		GetLabelByID(ctx context.Context, labelID models.LabelID) (models.Label, error)

		// GetLabelByName gets a label from the database
		GetLabelByName(ctx context.Context, name string) (models.Label, error)

		// UpdateLabel updates a label in the database
		UpdateLabel(ctx context.Context, label models.Label) error

		// DeleteLabel deletes a label from the database
		DeleteLabel(ctx context.Context, labelID models.LabelID) error

		// CreateLabel creates a label in the database
		CreateLabel(ctx context.Context, label models.Label) error

		// ListLabels lists all labels in the database
		ListLabels(ctx context.Context) ([]models.Label, error)

		// ListLinksByLabel lists all links with a label in the database
		ListLinksByLabel(ctx context.Context, labelID models.LabelID) ([]models.Link, error)

		// * Namespace
		// GetNamespace gets a Namespace from the database
		GetNamespace(ctx context.Context, namespace string) (models.Namespace, error)

		// DeleteNamespace deletes a Namespace from the database
		DeleteNamespace(ctx context.Context, namespace string) error

		// CreateNamespace creates a Namespace in the database
		CreateNamespace(ctx context.Context, Namespace models.NamespaceRequest) (models.Namespace, error)

		// ListNamespaces lists all Namespaces in the database
		ListNamespaces(ctx context.Context) ([]models.Namespace, error)

		// AddLinkToNamespace adds a link to a Namespace in the database
		// AddLinkToNamespace(ctx context.Context, namespace string, linkID models.LinkID) (error

		// RemoveLinkFromNamespace removes a link from a Namespace in the database
		// RemoveLinkFromNamespace(ctx context.Context, namespace string, linkID models.LinkID) error
	}
)
