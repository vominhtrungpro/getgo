package users

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kytruong0712/getgo/api/internal/controller/users"
	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHander_Users_GetByUId(t *testing.T) {
	type mockCtrl struct {
		expCall bool
		input   int64
		output  model.User
		err     error
	}

	tcs := map[string]struct {
		mockCtrl mockCtrl
		expBody  string
		expError error
		expCode  int
	}{
		"success": {
			expBody: `{"ID":15963452203860227,"ExternalID":"","Username":"vominhtrung1","Password":"123456789","Email":"vominhtrungpro1@gmail.com","Age":25,"RefreshToken":"","RefreshtokenExpiretime":"0001-01-01T00:00:00Z"}`,
			expCode: http.StatusOK,
			mockCtrl: mockCtrl{
				expCall: true,
				input:   15963571540197635,
				output: model.User{
					ID:       15963452203860227,
					Username: "vominhtrung1",
					Password: "123456789",
					Email:    "vominhtrungpro1@gmail.com",
					Age:      25,
				},
			},
		},
		//TODO adding other case
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			mockCtrl := new(users.MockController)
			if tc.mockCtrl.expCall {
				mockCtrl.On("GetById", mock.Anything, tc.mockCtrl.input).Return(tc.mockCtrl.output, tc.mockCtrl.err)
			}
			h := Handler{userCtrl: mockCtrl}
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("id", "15963571540197635")
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/users/id/{id}", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
			handler := http.HandlerFunc(h.GetByUId())
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
