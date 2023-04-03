package inventory

import (
	"database/sql"
	"context"
	"github.com/kytruong0712/getgo/api/internal/model"
)

type Repository interface {
	// specification
	CreateProduct(context.Context, model.Product) (model.Product, error)
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
