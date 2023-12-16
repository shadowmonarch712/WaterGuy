package handlers

import (
	"gofr.dev/pkg/gofr"
	"waterguy/models"
	"waterguy/stores"
)

type handler struct {
	store stores.User
}

func New(c stores.User) handler {
	return handler{store: c}
}

func (h handler) CreateUserEntry(ctx *gofr.Context) (interface{}, error) {
	var c models.UserEntry

	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}

	err = h.store.CreateUserEntry(ctx, c)
	if err != nil {
		return nil, err
	}

	return c.Value, nil
}
func (h handler) FetchUserEntry(ctx *gofr.Context) (interface{}, error) {
	params := ctx.Params()
	userID := params["userID"]
	resp, err := h.store.FetchUserEntry(ctx, userID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (h handler) UpdateUserEntry(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.UpdateUserEntry(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (h handler) DeleteUserEntry(ctx *gofr.Context) (interface{}, error) {
	userID := ctx.Params()["userID"]
	resp, err := h.store.DeleteUserEntry(ctx, userID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
