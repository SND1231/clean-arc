package router

import (
	"ddd/infrastructure/setting"
	"ddd/presenter/handler"

	"github.com/go-chi/chi"
)

func Get(settings setting.Setting) *chi.Mux {
	r := chi.NewRouter()

	auth := handler.NewAuthHandler(settings)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
		r.Post("/add", auth.Add)
	})
	return r
}
