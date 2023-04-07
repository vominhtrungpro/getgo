package products

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// Create creates new product
	Create(context.Context, CreateInput) (model.Product, error)
	// GetWithAssociateCategories gets single product with associate categories
	GetWithAssociateCategories(context.Context, string) (model.ProductWithCategories, error)
}

// New returns an implementation instance which satisfying Controller
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
