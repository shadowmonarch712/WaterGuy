package routes

import (
    "gofr.dev/pkg/gofr"
    "waterguy/handlers"
    "waterguy/stores/User"
)

func SetupRoutes(app *gofr.Gofr) {

    store := User.New()
	h := handlers.New(store)

    app.POST("/store", h.CreateUserEntry)
    app.GET("/fetch", h.FetchUserEntry)
    app.PUT("/update", h.UpdateUserEntry)
    app.DELETE("/delete", h.DeleteUserEntry)
}
