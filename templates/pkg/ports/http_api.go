package ports

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ServerInterface interface {
	//  (POST /v1/cancel-resource)
	CancelResource(w http.ResponseWriter, r *http.Request)
	//  (GET /v1/resources)
	GetAllResources(w http.ResponseWriter, r *http.Request)
	//  (GET /v1/resources/#{id})
	GetResource(w http.ResponseWriter, r *http.Request)
}

func InitializeRoutes(si ServerInterface, r chi.Router) chi.Router {
	r.Route("/v1/resources", func(r chi.Router) {
		r.Get("/", si.GetAllResources)
		r.Route("/{resourceID}", func(r chi.Router) {
			r.Get("/", si.GetResource)
		})
	})

	return r
}
