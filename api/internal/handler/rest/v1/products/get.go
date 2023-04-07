package products

import (
	"net/http"
	"strings"

	"github.com/kytruong0712/getgo/api/internal/httpserver"
	"github.com/go-chi/chi/v5"
	"log"
	"fmt"
)

// ProductWithCategoriesResponse ...
type ProductWithCategoriesResponse struct {
}

// GetWithAssociateCategories gets single product with associate categories by extID
func (h Handler) GetWithAssociateCategories() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("fired to Get...")
		log.Println("fired to Get...")

		ctx := r.Context()

		extID := chi.URLParam(r, "extID")
		if strings.TrimSpace(extID) == "" {
			return errInvalidExtID
		}

		log.Println("extID: ", extID)
		fmt.Println("extID: ", extID)

		resp, err := h.productCtrl.GetWithAssociateCategories(ctx, extID)
		fmt.Println("err: ", err)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)

		return nil
	})
}
