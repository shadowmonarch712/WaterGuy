package stores

import (
	"gofr.dev/pkg/gofr"
	"waterguy/models"
)

type User interface {
	CreateUserEntry(ctx *gofr.Context, model models.UserEntry) error
	FetchUserEntry(ctx *gofr.Context) (interface{}, error)
	UpdateUserEntry(ctx *gofr.Context) (interface{}, error)
	DeleteUserEntry(ctx *gofr.Context) (interface{}, error)
}