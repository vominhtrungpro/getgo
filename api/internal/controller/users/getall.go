package users

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
)

// Create creates new product
func (i impl) GetAll(ctx context.Context) ([]model.User, error) {
	return i.repo.Inventory().GetAllUser(ctx)
}
