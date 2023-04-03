package products

import (
	"net/http"

	"github.com/kytruong0712/getgo/api/internal/httpserver"
)

const ErrCodeValidationFailed = "validation_failed"

var (
	errInvalidName        = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Name"}
	errInvalidDescription = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Description"}
	errInvalidPrice       = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Price"}
)
