package authentication

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	Login(context.Context, LoginInput) (TokenOutput, string, error)
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
