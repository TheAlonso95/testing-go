package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/goschool/crud/api"
)

func SetupRoutes(userHandler api.UserHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", api.HandleHelloWorld)
	router.Post("/echo", api.HandleEchoUser)
	router.Post("/register", userHandler.HandlerRegisterUser)
	router.Post("/login", userHandler.HandlerLoginUser)
	return router
}
