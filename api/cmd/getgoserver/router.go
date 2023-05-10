package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/kytruong0712/getgo/api/graph"
	authen "github.com/kytruong0712/getgo/api/internal/controller/authentication"
	product "github.com/kytruong0712/getgo/api/internal/controller/products"
	user "github.com/kytruong0712/getgo/api/internal/controller/users"
	"github.com/kytruong0712/getgo/api/internal/handler/rest/v1/authentication"
	"github.com/kytruong0712/getgo/api/internal/handler/rest/v1/products"
	"github.com/kytruong0712/getgo/api/internal/handler/rest/v1/users"
)

type router struct {
	productCtrl product.Controller
	userCtrl    user.Controller
	authenCtrl  authen.Controller
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
			pattern := prefix + "/products"

			r.Post(pattern, products.New(rtr.productCtrl).Create())
		})

		r.Group(func(r chi.Router) {
			pattern := prefix + "/users"

			r.Get(pattern+"/username/{username}", users.New(rtr.userCtrl).GetByUsername())
			r.Get(pattern+"/id/{id}", users.New(rtr.userCtrl).GetByUId())
			r.Get(pattern, users.New(rtr.userCtrl).GetAll())
			r.Post(pattern, users.New(rtr.userCtrl).Create())
			r.Put(pattern, users.New(rtr.userCtrl).Update())
			r.Delete(pattern+"/id/{id}", users.New(rtr.userCtrl).DeleteById())
			r.Post(pattern+"/import", users.New(rtr.userCtrl).Import())
		})

		r.Group(func(r chi.Router) {
			pattern := prefix + "/login"

			r.Post(pattern, authentication.New(rtr.authenCtrl).Login())
		})

	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserCtrl: rtr.userCtrl}}))
	r.Handle("/graphql", srv)
}
