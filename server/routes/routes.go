package routes

import (
    "waterguy/controllers"
    "gofr.dev/pkg/gofr"
)

func SetupRoutes(app *gofr.Gofr) {
    app.POST("/store", controllers.CreateUserEntry)
    app.GET("/fetch", controllers.FetchUserEntry)
    app.PUT("/update", controllers.UpdateUserEntry)
    app.DELETE("/delete", controllers.DeleteUserEntry)
}
