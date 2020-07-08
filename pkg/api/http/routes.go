package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	errs "github.com/pkg/errors"
)

// routes returns a new chi mux with routes and handlers defined
func routes(server Server) *chi.Mux {
	r := chi.NewRouter()
	cors := corsOptions()

	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.RequestID,
			middleware.RealIP,
			middleware.Logger,
			middleware.Recoverer,
			middleware.StripSlashes,
			render.SetContentType(render.ContentTypeJSON),
		)
		r.Route("/v1/api", func(r chi.Router) {
			// Place server routes here ex.
			// r.Get("/routeName", server.HandleRouteName())
			r.Get("/public", server.HandleGetPublic())
			r.Get("/inbox", server.HandleGetInbox())
			r.Get("/outbox", server.HandleGetOutbox())
			r.Post("/public", server.HandleSendPublic())
			r.Post("/private", server.HandleSendPrivate())
			r.Get("/users", server.HandleGetUsers())
		})
	})

	return r
}

// Walk walks through and prints out the routes and their http methods
func walk(r *chi.Mux) error {
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Printf("logging err: %s\n", err.Error())
		return errs.Wrap(err, "api.Routes.Walk")
	}

	return nil
}

func walkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	log.Printf("%s %s \n", method, route)
	return nil
}

func corsOptions() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	})

	return cors
}
