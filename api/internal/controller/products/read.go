package products

import (
	"context"
	"errors"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository/inventory"
)

// GetWithAssociateCategories gets single product with associate categories
func (i impl) GetWithAssociateCategories(ctx context.Context, extID string) (model.ProductWithCategories, error) {
	m, err := i.repo.Inventory().GetProductWithCategories(ctx, extID)
	if err != nil {
		if errors.Is(err, inventory.ErrNotFound) {
			return model.ProductWithCategories{}, ErrProductNotFound
		}

		return model.ProductWithCategories{}, err
	}

	return m, nil
}
