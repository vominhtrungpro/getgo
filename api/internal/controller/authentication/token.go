package authentication

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kytruong0712/getgo/api/internal/repository/dbmodel"
)

type TokenOutput struct {
	AccessToken  string
	RefreshToken string
}

func CreateAccessToken(user *dbmodel.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id":    user.ExternalID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 240).Unix(),
		"role":  "None",
		"name":  user.Username,
	})

	tokenString, err := token.SignedString([]byte("SecretYouShouldHide"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	return token, nil
}
