package users

import (
	"net/http"

	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product

type ImportInput struct {
	Url string
}

func (h Handler) Import() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		file, _, err := r.FormFile("file")
		if err != nil {
			return err
		}

		resp, err := h.userCtrl.Import(ctx, file)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)
		return nil
	})
}
