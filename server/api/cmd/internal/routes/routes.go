package routes

import (
	"api/cmd/internal/auth"
	"api/cmd/internal/handlers"
	"api/cmd/internal/middlewares"
	"api/cmd/internal/postgresrepo"
	"api/cmd/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(queries *postgresrepo.Queries) http.Handler {
	r := chi.NewRouter()

	services := services.NewService(queries)
	JTWConfig := auth.NewJWTConfig()

	handler := handlers.NewHandler(services, JTWConfig)

	r.With(middlewares.Cors)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/", handler.IndexHandler)
	r.Post("/register", handler.CreateUser)
	r.Post("/login", handler.Login)
	r.Post("/forget-password", handler.ForgotPassword)
	r.Post("/recover-password", handler.RecoverPassword)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Get("/user", handler.GetUser)
		r.Post("/logout", handler.Logout)

		r.Post("/import", handler.ImportFile)
		r.Get("/customers", handler.GetAllCustomers)
		r.Get("/resources", handler.GetAllResources)
		r.Get("/categories", handler.GetAllCategories)

	})

	return r
}
