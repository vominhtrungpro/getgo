package products

import "github.com/kytruong0712/getgo/api/internal/controller/products"

// Handler is the web handler for this pkg
type Handler struct {
	productCtrl products.Controller
}

// New instantiates a new Handler and returns it
func New(productCtrl products.Controller) Handler {
	return Handler{
		productCtrl: productCtrl,
	}
}
