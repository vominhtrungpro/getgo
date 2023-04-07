package inventory

import (
	"context"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository/dbmodel"
	"github.com/kytruong0712/getgo/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"errors"
	"database/sql"
	"time"
)

// CreateProduct inserts a record to product table
func (i impl) CreateProduct(ctx context.Context, m model.Product) (model.Product, error) {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return model.Product{}, err
	}

	o := dbmodel.Product{
		ID:          id,
		ExternalID:  m.ExternalID,
		Name:        m.Name,
		Description: m.Description,
		Status:      m.Status.String(),
		Price:       m.Price,
	}

	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return model.Product{}, pkgerrors.WithStack(err)
	}

	m.ID = id
	m.CreatedAt = o.CreatedAt
	m.UpdatedAt = o.UpdatedAt

	return m, nil
}

// GetProductWithCategories retrieves specific product with list of categories by extID
func (i impl) GetProductWithCategories(ctx context.Context, productExternalID string) (model.ProductWithCategories, error) {
	dbm, err := dbmodel.Products(
		qm.Load("ProductCategories"),
		qm.Load("ProductCategories.Category"),
		dbmodel.ProductWhere.ExternalID.EQ(productExternalID),
	).One(ctx, i.db)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ProductWithCategories{}, ErrNotFound
		}

		return model.ProductWithCategories{}, pkgerrors.WithStack(err)
	}

	var m model.ProductWithCategories
	m.ID = dbm.ID
	m.ExternalID = dbm.ExternalID
	m.Price = dbm.Price
	m.Name = dbm.Name
	m.Description = dbm.Description
	m.Status = model.ProductStatus(dbm.Status)
	m.CreatedAt = dbm.CreatedAt
	m.UpdatedAt = dbm.UpdatedAt

	if dbm.R != nil && dbm.R.ProductCategories != nil {
		for _, item := range dbm.R.ProductCategories {
			if item.R != nil && item.R.Category != nil {
				m.Categories = append(m.Categories, model.Category{
					ID:          item.R.Category.ID,
					ExternalID:  item.R.Category.ExternalID,
					Description: item.R.Category.Description,
					Name:        item.R.Category.Name,
					Status:      model.CategoryStatus(item.R.Category.Status),
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				})
			}
		}
	}

	return m, nil
}
