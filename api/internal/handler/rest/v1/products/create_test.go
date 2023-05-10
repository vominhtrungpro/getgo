package products

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kytruong0712/getgo/api/internal/controller/products"
	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Products_Create(t *testing.T) {
	type mockCtrl struct {
		expCall bool
		input   products.CreateInput
		output  model.Product
		err     error
	}

	tcs := map[string]struct {
		givenInput string
		mockCtrl   mockCtrl
		expBody    string
		expError   error
		expCode    int
	}{
		"success": {
			givenInput: `{"name":"Mackbook Pro M1","description":"A Mackbook Pro","price":1500}`,
			expBody:    `{"ID":101,"Price":1500,"ExternalID":"","Description":"A Mackbook Pro","Name":"Mackbook Pro M1","Status":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`,
			expCode:    http.StatusOK,
			mockCtrl: mockCtrl{
				expCall: true,
				input: products.CreateInput{
					Name:        "Mackbook Pro M1",
					Description: "A Mackbook Pro",
					Price:       1500,
				},
				output: model.Product{
					ID:          101,
					Name:        "Mackbook Pro M1",
					Description: "A Mackbook Pro",
					Price:       1500,
				},
			},
		},
		"error when invalid products name": {
			givenInput: `{"description":"A Mackbook Pro","price":1500}`,
			expBody:    `{"error":"validation_failed","error_description":"Invalid Product Name"}`,
			expCode:    http.StatusBadRequest,
			expError:   errInvalidName,
		},
		//TODO adding other case
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(tc.givenInput))

			mockCtrl := new(products.MockController)
			if tc.mockCtrl.expCall {
				mockCtrl.On("Create", mock.Anything, tc.mockCtrl.input).Return(tc.mockCtrl.output, tc.mockCtrl.err)

			}

			h := Handler{productCtrl: mockCtrl}
			handler := http.HandlerFunc(h.Create())
			handler.ServeHTTP(res, req)
			mockCtrl.AssertExpectations(t)
			if tc.expError != nil {
				require.Equal(t, res.Code, tc.expCode)
				require.Equal(t, tc.expBody, res.Body.String())
			} else {
				require.NoError(t, tc.expError)
				require.Equal(t, tc.expCode, res.Code)
				require.Equal(t, tc.expBody, res.Body.String())
			}
		})
	}
}
