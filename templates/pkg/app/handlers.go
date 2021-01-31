package app

import (
	"net/http"
)

func (s *server) createResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.service.Create()
		w.Write([]byte("createResource"))
	})
}

func (s *server) getResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.service.Get()
		w.Write([]byte("getResource"))
	})
}

func (s *server) getAllResources() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.service.GetAll()
		w.Write([]byte("getAllResources"))
	})
}

func (s *server) updateResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.service.Update()
		w.Write([]byte("updateResource"))
	})
}

func (s *server) deleteResource() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.service.Delete()
		w.Write([]byte("deleteResource"))
	})
}
