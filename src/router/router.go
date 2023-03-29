package router

import (
	"ddd/infrastructure/setting"
	worker "ddd/services/worker/presenter"

	"github.com/go-chi/chi"
)

func Get(settings setting.Setting) *chi.Mux {
	r := chi.NewRouter()

	workerHandler := worker.NewAuthHandler(settings)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", workerHandler.Login)
		r.Post("/add", workerHandler.Add)
	})
	return r
}
