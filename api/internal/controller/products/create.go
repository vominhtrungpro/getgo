package products

import (
	"context"
	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/pkg/uid"
)

// CreateInput represents for input to create Product
type CreateInput struct {
	Name        string
	Description string
	Price       int64
}

// Create creates new product
func (i impl) Create(ctx context.Context, input CreateInput) (model.Product, error) {
	extID, err := uid.Generate()
	if err != nil {
		return model.Product{}, nil
	}

	return i.repo.Inventory().CreateProduct(ctx, model.Product{
		ExternalID:  extID,
		Price:       input.Price,
		Description: input.Description,
		Name:        input.Name,
		Status:      model.ProductStatusActive,
	})
}
