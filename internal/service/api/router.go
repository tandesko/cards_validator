package api

import (
	"github.com/go-chi/chi"
	"github.com/tandesko/cards_validator/internal/service/api/handlers"
	"github.com/tandesko/cards_validator/internal/service/api/helpers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log)),
	)

	r.Route("/integrations/cards", func(r chi.Router) {
		r.Route("/verify-card", func(r chi.Router) {
			r.Post("/", handlers.ValidateCard)
		})
	})

	return r
}
