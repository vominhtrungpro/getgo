package users

import (
	"context"
)

// Create creates new product
func (i impl) DeleteById(ctx context.Context, id int64) error {
	return i.repo.Inventory().DeleteUserById(ctx, id)
}
