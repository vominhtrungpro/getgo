package authentication

import (
	"context"
)

// CreateInput represents for input to create Product
type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	AccessToken string
}

// Create creates new product
func (i impl) Login(ctx context.Context, input LoginInput) (TokenOutput, string, error) {
	var token TokenOutput
	userdb, message, err := i.repo.Inventory().CheckUsernameAndPassword(ctx, input.Username, input.Password)
	if err != nil {
		return token, message, nil
	}
	accessToken, errr := CreateAccessToken(userdb)
	if errr != nil {
		return token, "Cant create accesstoken", errr
	}
	refreshToken, errrr := CreateRefreshToken()
	if errrr != nil {
		return token, "Cant create refreshToken", errrr
	}
	errrrr := i.repo.Inventory().UpdateToken(ctx, userdb, refreshToken)
	if errrrr != nil {
		return token, "Cant update refreshToken", errrrr
	}
	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return token, message, nil
}
