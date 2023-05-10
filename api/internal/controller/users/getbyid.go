package users

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
)

// Create creates new product
func (i impl) GetById(ctx context.Context, id int64) (model.User, error) {
	return i.repo.Inventory().GetUserById(ctx, id)
}
