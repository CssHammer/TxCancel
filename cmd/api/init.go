package main

import (
	"github.com/go-chi/chi"

	"txCancel/http"
	"txCancel/middleware"
)

func getRouter(handler *http.Handler, middlewares *middleware.Middlewares) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		middlewares.SourceType,
		middlewares.JSON,
	)

	r.Get("/user", handler.GetUser)
	r.Post("/tx", handler.PostTx)

	return r
}
