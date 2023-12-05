package api

import (
	"mamuro_mail-api/api/controllers"
	"net/http"
	"os"

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

func static(r *chi.Mux) {
	root := "./view/dist"
	fs := http.FileServer(http.Dir(root))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	})
}

func RunServer(port int) {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	addCors(r)
	static(r)

	apiv1 := chi.NewRouter()

	controllers.AddControllers(apiv1)
	r.Mount("/api/v1", apiv1)

	http.ListenAndServe(":3000", r)
}
