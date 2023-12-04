package api

import (
	"mamuro_mail-api/api/controllers"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func addCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))
}

func RunServer(port int) {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json","text/xml"))

	addCors(r)
	// static(r)

	apiv1 := chi.NewRouter()
	
	controllers.AddControllers(apiv1)
	r.Mount("/api/v1", apiv1)

	http.ListenAndServe(":3000", r)
}