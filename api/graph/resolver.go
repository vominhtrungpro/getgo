package graph

import (
	"github.com/kytruong0712/getgo/api/graph/model"
	"github.com/kytruong0712/getgo/api/internal/controller/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]model.Character
	CharacterStory []*model.Character
	UserCtrl       users.Controller
}
