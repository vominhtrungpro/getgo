package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kytruong0712/getgo/api/internal/controller/users"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product
func (h Handler) Update() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := users.UpdateInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		if err := validateupdate(input); err != nil {
			return err
		}

		err := h.userCtrl.Update(ctx, input)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, "Success")

		return nil
	})
}

func validateupdate(input users.UpdateInput) error {
	if strings.TrimSpace(input.Username) == "" {
		return errInvalidUsername
	}

	if strings.TrimSpace(input.Password) == "" {
		return errInvalidPassword
	}

	if strings.TrimSpace(input.Email) == "" {
		return errInvalidEmail
	}

	if input.Age <= 0 {
		return errInvalidAge
	}

	return nil
}
