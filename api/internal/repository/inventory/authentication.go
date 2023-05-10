package inventory

import (
	"context"
	"fmt"

	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository/dbmodel"
	pkgerrors "github.com/pkg/errors"
)

func (i impl) CheckUsernameAndPassword(ctx context.Context, username string, password string) (*dbmodel.User, string, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(username)).One(ctx, i.db)
	var user model.User
	if userdb == nil {
		return userdb, "User not found!", pkgerrors.Wrap(err, "User not found!")
	}
	if err != nil {
		return userdb, "Something went wrong!", pkgerrors.WithStack(err)
	}
	if userdb.Password != password {
		err := fmt.Errorf("error message with value ")
		return userdb, "Password not match!", err
	}
	user.ID = userdb.ID
	user.ExternalID = userdb.ExternalID
	user.Username = userdb.Username
	user.Password = userdb.Password
	user.Email = userdb.Password
	user.Age = userdb.Age
	return userdb, "Login Success", nil
}
