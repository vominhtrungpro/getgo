package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kytruong0712/getgo/api/internal/controller/users"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

// Create creates new product
func (h Handler) Create() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := users.CreateInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		if err := validate(input); err != nil {
			return err
		}

		resp, err := h.userCtrl.Create(ctx, input)
		if err != nil {
			return err
		}

		httpserver.RespondJSON(w, resp)
		return nil
	})
}

func validate(input users.CreateInput) error {
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

// func validateAll(input users.CreateInput) []error {
// 	var errs []error
// 	if strings.TrimSpace(input.Username) == "" {
// 		errs = append(errs, errInvalidUsername)
// 	}

// 	if strings.TrimSpace(input.Password) == "" {
// 		errs = append(errs, errInvalidPassword)
// 	}

// 	if strings.TrimSpace(input.Email) == "" {
// 		errs = append(errs, errInvalidEmail)
// 	}

// 	if input.Age <= 0 {
// 		errs = append(errs, errInvalidAge)
// 	}

// 	if len(errs) > 0 {
// 		return errs
// 	}

// 	return nil
// }
