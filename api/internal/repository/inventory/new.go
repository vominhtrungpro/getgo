package inventory

import (
	"context"
	"database/sql"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository/dbmodel"
)

type Repository interface {
	// specification
	CreateProduct(context.Context, model.Product) (model.Product, error)
	CreateUser(context.Context, model.User) (model.User, error)
	GetAllUser(context.Context) ([]model.User, error)
	GetUserByUsername(context.Context, string) (model.User, error)
	GetUserById(context.Context, int64) (model.User, error)
	DeleteUserById(context.Context, int64) error
	UpdateUserById(context.Context, model.User) error
	CheckUsernameAndPassword(context.Context, string, string) (*dbmodel.User, string, error)
	UpdateToken(context.Context, *dbmodel.User, string) error
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}

type impl struct {
	db *sql.DB
}
