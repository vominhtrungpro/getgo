package users

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/pkg/uid"
)

// CreateInput represents for input to create Product
type CreateInput struct {
	Username string
	Password string
	Email    string
	Age      int64
}

// Create creates new product
func (i impl) Create(ctx context.Context, input CreateInput) (model.User, error) {
	extID, err := uid.Generate()
	if err != nil {
		return model.User{}, nil
	}

	return i.repo.Inventory().CreateUser(ctx, model.User{
		ExternalID: extID,
		Username:   input.Username,
		Password:   input.Password,
		Email:      input.Email,
		Age:        input.Age,
	})
}
