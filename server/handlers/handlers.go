package handlers

import (
	"waterguy/stores"
	"waterguy/models"
	"gofr.dev/pkg/gofr"
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
	resp, err := h.store.FetchUserEntry(ctx)
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
	resp, err := h.store.DeleteUserEntry(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}