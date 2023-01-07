package router

import (
	"clean-arc/presenter/handler"

	"github.com/go-chi/chi"
)

func Get() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetUsers)
	})
	return r
}
