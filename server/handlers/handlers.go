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

// func (h handler) CreateUserEntry(ctx *gofr.Context, c models.UserEntry) (interface{}, error) {
// 	// userID := ctx.Param("userID")
// 	resp, err := h.store.CreateUserEntry(ctx, c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }
func (h handler) CreateUserEntry(ctx *gofr.Context) (interface{}, error) {
	// resp, err := h.store.CreateUserEntry(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return resp, nil
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
	// userID := ctx.Param("userID")
	resp, err := h.store.FetchUserEntry(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (h handler) UpdateUserEntry(ctx *gofr.Context) (interface{}, error) {
	// userID := ctx.Param("userID")
	resp, err := h.store.UpdateUserEntry(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (h handler) DeleteUserEntry(ctx *gofr.Context) (interface{}, error) {
	// userID := ctx.Param("userID")
	resp, err := h.store.DeleteUserEntry(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
// func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
// 	userID := ctx.Param("userID")
// 	resp, err := h.store.Get(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

// func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
// 	var c models.User

// 	err := ctx.Bind(&c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = h.store.Create(ctx, c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return "New User Added!", nil
// }

// func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
// 	userID := ctx.Param("userID")

// 	deleteCount, err := h.store.Delete(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return fmt.Sprintf("%v Users Deleted!", deleteCount), nil
// }