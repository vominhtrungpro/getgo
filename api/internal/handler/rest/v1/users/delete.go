package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product
func (h Handler) DeleteById() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}
		errr := h.userCtrl.DeleteById(ctx, i)
		if errr != nil {
			return errr
		}
		httpserver.RespondJSON(w, "Success")
		return nil
	})
}
