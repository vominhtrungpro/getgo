package inventory

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kytruong0712/getgo/api/internal/httpserver"
	"github.com/kytruong0712/getgo/api/internal/model"
	"github.com/kytruong0712/getgo/api/internal/repository/dbmodel"
	"github.com/kytruong0712/getgo/api/internal/repository/generator"
	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MyError struct {
	Err    error
	Reason string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s: %s", e.Err.Error(), e.Reason)
}

func (i impl) CreateUser(ctx context.Context, m model.User) (model.User, error) {
	id, err := generator.ProductSNF.Generate()
	if err != nil {
		return model.User{}, err
	}

	userdb, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(m.Username)).One(ctx, i.db)
	// if err != nil {
	// 	return model.User{}, pkgerrors.WithStack(err)
	// }
	if userdb != nil {
		return model.User{}, &httpserver.Error{Status: http.StatusBadRequest, Code: "validation_failed", Desc: "user exits"}
	}

	o := dbmodel.User{
		ID:         id,
		ExternalID: m.ExternalID,
		Username:   m.Username,
		Password:   m.Password,
		Email:      m.Email,
		Age:        m.Age,
	}

	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	m.ID = id
	return m, nil
}

func (i impl) GetAllUser(ctx context.Context) ([]model.User, error) {
	usersdb, err := dbmodel.Users().All(ctx, i.db)
	if err != nil {
		return []model.User{}, pkgerrors.WithStack(err)
	}
	var users []model.User
	for _, item := range usersdb {
		var user model.User
		user.ID = item.ID
		user.ExternalID = item.ExternalID
		user.Username = item.Username
		user.Password = item.Password
		user.Email = item.Email
		user.Age = item.Age
		user.RefreshToken = item.Refreshtoken.String
		user.RefreshtokenExpiretime = item.Refreshtokenexpiretime.Time
		users = append(users, user)
	}
	return users, nil
}

func (i impl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.Username.EQ(username)).One(ctx, i.db)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}
	var user model.User
	user.ID = userdb.ID
	user.ExternalID = userdb.ExternalID
	user.Username = userdb.Username
	user.Password = userdb.Password
	user.Email = userdb.Email
	user.Age = userdb.Age
	user.RefreshToken = userdb.Refreshtoken.String
	user.RefreshtokenExpiretime = userdb.Refreshtokenexpiretime.Time
	return user, nil
}

func (i impl) GetUserById(ctx context.Context, id int64) (model.User, error) {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(id)).One(ctx, i.db)
	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}
	var user model.User
	user.ID = userdb.ID
	user.ExternalID = userdb.ExternalID
	user.Username = userdb.Username
	user.Password = userdb.Password
	user.Email = userdb.Email
	user.Age = userdb.Age
	user.RefreshToken = userdb.Refreshtoken.String
	user.RefreshtokenExpiretime = userdb.Refreshtokenexpiretime.Time
	return user, nil
}

func (i impl) UpdateUserById(ctx context.Context, m model.User) error {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(m.ID)).One(ctx, i.db)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	userdb.ExternalID = m.ExternalID
	userdb.Username = m.Username
	userdb.Password = m.Password
	userdb.Age = m.Age
	userdb.Email = m.Email
	_, errr := userdb.Update(ctx, i.db, boil.Infer())
	if errr != nil {
		return pkgerrors.WithStack(err)
	}
	return nil
}

func (i impl) DeleteUserById(ctx context.Context, id int64) error {
	userdb, err := dbmodel.Users(dbmodel.UserWhere.ID.EQ(id)).One(ctx, i.db)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	_, errr := userdb.Delete(ctx, i.db)
	if errr != nil {
		return pkgerrors.WithStack(errr)
	}
	return nil
}

func (i impl) UpdateToken(ctx context.Context, userdb *dbmodel.User, token string) error {
	userdb.Refreshtoken = null.String{String: token, Valid: true}
	expdate := time.Now()
	userdb.Refreshtokenexpiretime = null.Time{Time: expdate, Valid: true}
	_, errr := userdb.Update(ctx, i.db, boil.Infer())
	if errr != nil {
		return pkgerrors.WithStack(errr)
	}
	return nil
}
