package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *server) createResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resource := &ResourcePayload{}
		err := s.payloadSerializer(r, resource)
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		createdResource, err := s.Service.Create(ctx, resource.toModel())
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		s.successSerializer(w, createdResource)
	})
}

func (s *server) getResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resourceID, err := stringToUint64(chi.URLParam(r, "resourceID"))
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		resource, err := s.Service.Get(ctx, resourceID)
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		s.successSerializer(w, resource)
	})
}

func (s *server) getAllResources() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resources, err := s.Service.GetAll(ctx)
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		s.successSerializer(w, resources)
	})
}

func (s *server) updateResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resource := &ResourcePayload{}
		err := s.payloadSerializer(r, resource)
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		err = s.Service.Update(ctx, resource.toModel())
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (s *server) deleteResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resourceID, err := stringToUint64(chi.URLParam(r, "resourceID"))
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		err = s.Service.Delete(ctx, resourceID)
		if err != nil {
			s.errorSerializer(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
