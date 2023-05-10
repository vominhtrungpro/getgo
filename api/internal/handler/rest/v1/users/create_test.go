package users

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kytruong0712/getgo/api/internal/controller/users"
	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHander_Users_Create(t *testing.T) {
	type mockCtrl struct {
		expCall bool
		input   users.CreateInput
		output  model.User
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
			givenInput: `{"username":"vominhtrungpro","password":"123456789","email":"vominhtrungpro@gmail.com","age":25}`,
			expBody:    `{"ID":1,"ExternalID":"","Username":"vominhtrungpro","Password":"123456789","Email":"vominhtrungpro@gmail.com","Age":25,"RefreshToken":"","RefreshtokenExpiretime":"0001-01-01T00:00:00Z"}`,
			expCode:    http.StatusOK,
			mockCtrl: mockCtrl{
				expCall: true,
				input: users.CreateInput{
					Username: "vominhtrungpro",
					Password: "123456789",
					Email:    "vominhtrungpro@gmail.com",
					Age:      25,
				},
				output: model.User{
					ID:       1,
					Username: "vominhtrungpro",
					Password: "123456789",
					Email:    "vominhtrungpro@gmail.com",
					Age:      25,
				},
			},
		},
		"error when invalid users username": {
			givenInput: `{"password":"123456789","email":"vominhtrungpro@gmail.com","age":25}`,
			expBody:    `{"error":"validation_failed","error_description":"Invalid User Username"}`,
			expCode:    http.StatusBadRequest,
			expError:   errInvalidUsername,
		},
		//TODO adding other case
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.givenInput))

			mockCtrl := new(users.MockController)
			if tc.mockCtrl.expCall {
				mockCtrl.On("Create", mock.Anything, tc.mockCtrl.input).Return(tc.mockCtrl.output, tc.mockCtrl.err)
			}

			h := Handler{userCtrl: mockCtrl}
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
