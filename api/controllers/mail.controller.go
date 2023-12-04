package controllers

import (
	"mamuro_mail-api/api/services"

	"github.com/go-chi/chi/v5"
)

func MailController(r *chi.Mux) {
	r.Route("/mail", func(r chi.Router) {
		r.Post("/", services.SearchData)
	})
}