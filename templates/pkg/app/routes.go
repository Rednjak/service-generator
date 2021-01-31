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
	s.router.Route("/v1/resources", func(chi.Router) {
		s.router.Get("/", s.getAllResources())
		s.router.Route("/{resourceID}", func(chi.Router) {
			s.router.Post("/", s.createResource())
			s.router.Get("/", s.getResource())
			s.router.Put("/", s.updateResource())
			s.router.Delete("/", s.deleteResource())
		})
	})
}
