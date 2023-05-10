package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/kytruong0712/getgo/api/internal/controller/authentication"
	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

type Response struct {
	Message string                     `json:"message"`
	Token   authentication.TokenOutput `json:"token"`
}

func (h Handler) Login() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		input := authentication.LoginInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			return err
		}

		tokenresp, messresp, err := h.authenCtrl.Login(ctx, input)
		if err != nil {
			httpserver.RespondJSON(w, err)
			return err
		}
		var resp Response
		resp.Message = messresp
		resp.Token = tokenresp

		httpserver.RespondJSON(w, resp)
		return nil
	})
}
