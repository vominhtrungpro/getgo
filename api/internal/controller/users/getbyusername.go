package users

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
)

// Create creates new product
func (i impl) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return i.repo.Inventory().GetUserByUsername(ctx, username)
}
