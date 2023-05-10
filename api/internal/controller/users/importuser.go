package users

import (
	"context"
	"encoding/csv"
	"mime/multipart"
	"strconv"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/pkg/uid"
)

type ImportUserInput struct {
	Username string
	Password string
	Email    string
	Age      string
}

func (i impl) Import(ctx context.Context, file multipart.File) ([]model.User, error) {
	var users []model.User

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return users, err
	}
	for _, line := range csvLines {
		user := ImportUserInput{
			Username: line[0],
			Password: line[1],
			Email:    line[2],
			Age:      line[3],
		}

		age, err := strconv.Atoi(user.Age)
		if err != nil {
			return users, err
		}

		extID, err := uid.Generate()
		if err != nil {
			return users, err
		}

		userdb, err := i.repo.Inventory().CreateUser(ctx, model.User{
			ExternalID: extID,
			Username:   user.Username,
			Password:   user.Password,
			Email:      user.Email,
			Age:        int64(age),
		})
		if err != nil {
			return users, nil
		}
		users = append(users, userdb)

	}
	return users, nil
}
