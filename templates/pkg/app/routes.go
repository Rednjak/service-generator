package app

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *server) routes() {
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(10 * time.Second))

	s.router.Route("/v1/resources", func(r chi.Router) {
		r.Get("/", s.getAllResources())
		r.Post("/", s.createResource())
		r.Route("/{resourceID}", func(r chi.Router) {
			r.Get("/", s.getResource())
			r.Put("/", s.updateResource())
			r.Delete("/", s.deleteResource())
		})
	})
}
