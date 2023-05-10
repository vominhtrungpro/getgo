package authentication

import "github.com/kytruong0712/getgo/api/internal/controller/authentication"

// Handler is the web handler for this pkg
type Handler struct {
	authenCtrl authentication.Controller
}

// New instantiates a new Handler and returns it
func New(authenCtrl authentication.Controller) Handler {
	return Handler{
		authenCtrl: authenCtrl,
	}
}
