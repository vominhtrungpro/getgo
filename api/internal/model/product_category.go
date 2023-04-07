package model

// ProductWithCategories represents the business model of product with list of categories
type ProductWithCategories struct {
	Product
	Categories []Category
}
