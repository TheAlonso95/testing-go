package main

import (
	"fmt"
	"net/http"

	"github.com/goschool/crud/api"
	"github.com/goschool/crud/db"

	"github.com/goschool/crud/routes"
)

func main() {
	fmt.Println("Program starting...")

	dbInstance, err := db.Open()
	if err != nil {
		panic(err)
	}

	userStore := db.NewSQLiteUserStore(dbInstance)
	userHandler := api.NewUserHandler(userStore)
	router := routes.SetupRoutes(*userHandler)

	//Start the app here:
	http.ListenAndServe(":8081", router)
}
