package main

import (
    "gofr.dev/pkg/gofr"
    "waterguy/database"
    "waterguy/routes"
)

type UserEntry struct {
    UserID string `json:"userID"`
    Value  int    `json:"value"`
}

func main() {
    // initialise gofr object
    app := gofr.New()

    // Initialize the database
    database.InitDB()
    // Setup routes
    routes.SetupRoutes(app)

    // app will start on port 8000
    app.Start()

    // connection close
    // defer func() {
    //     if err = client.Disconnect(context.TODO()); err != nil {
    //         log.Fatal(err)
    //     }
    // }()
}
 