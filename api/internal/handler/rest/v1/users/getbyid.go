package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product
func (h Handler) GetByUId() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		resp, err := h.userCtrl.GetById(ctx, i)
		if err != nil {
			return err
		}
		httpserver.RespondJSON(w, resp)

		return nil
	})
}
