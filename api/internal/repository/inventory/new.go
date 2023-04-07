package inventory

import (
	"database/sql"
	"context"
	"github.com/kytruong0712/getgo/api/internal/model"
)

type Repository interface {
	// CreateProduct inserts a record to product table
	CreateProduct(context.Context, model.Product) (model.Product, error)
	// GetProductWithCategories retrieves specific product with list of categories by extID
	GetProductWithCategories(context.Context, string) (model.ProductWithCategories, error)
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
