package router

import (
	"net/http"

	"secure-login/handlers"
	"secure-login/middleware"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handlers.Register)
	mux.HandleFunc("POST /login", handlers.Login)
	mux.HandleFunc("POST /logout", handlers.Logout)

	mux.Handle(
		"GET /dashboard",
		middleware.Auth(http.HandlerFunc(handlers.Dashboard)),
	)

	return mux
}
