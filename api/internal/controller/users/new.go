package users

import (
	"context"
	"mime/multipart"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	Create(context.Context, CreateInput) (model.User, error)
	GetAll(context.Context) ([]model.User, error)
	GetByUsername(context.Context, string) (model.User, error)
	GetById(context.Context, int64) (model.User, error)
	DeleteById(context.Context, int64) error
	Update(context.Context, UpdateInput) error
	Import(context.Context, multipart.File) ([]model.User, error)
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
