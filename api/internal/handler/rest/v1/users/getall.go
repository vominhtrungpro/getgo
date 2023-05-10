package users

import (
	"net/http"

	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product
func (h Handler) GetAll() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		resp, err := h.userCtrl.GetAll(ctx)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
