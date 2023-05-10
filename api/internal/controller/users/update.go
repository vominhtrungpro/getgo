package users

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
)

// CreateInput represents for input to create Product
type UpdateInput struct {
	ID         int64
	ExternalID string
	Username   string
	Password   string
	Email      string
	Age        int64
}

// Create creates new product
func (i impl) Update(ctx context.Context, input UpdateInput) error {
	return i.repo.Inventory().UpdateUserById(ctx, model.User{
		ID:         input.ID,
		ExternalID: input.ExternalID,
		Username:   input.Username,
		Password:   input.Password,
		Email:      input.Email,
		Age:        input.Age,
	})
}
