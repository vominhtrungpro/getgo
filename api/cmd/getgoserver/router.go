package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	product "github.com/kytruong0712/getgo/api/internal/controller/products"
	"github.com/kytruong0712/getgo/api/internal/handler/rest/v1/products"
)

type router struct {
	productCtrl product.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!!"))
		})
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		// products
		r.Group(func(r chi.Router) {
			pattern := prefix + "/products/"

			r.Post(pattern, products.New(rtr.productCtrl).Create())
		})
	})
}
