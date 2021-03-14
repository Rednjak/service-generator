package app

import (
	"log"
	"net/http"

	"github.com/repo/package/pkg/service"

	"github.com/go-chi/chi"
)

type server struct {
	router  *chi.Mux
	Service service.Service
}

// NewServer will initialize all the dependencies and return a server struct which can be used to start the server
func NewServer() server {
	s := server{}
	s.router = chi.NewRouter()
	s.routes()
	return s
}

// StartServer starts the server
func (s *server) StartServer() {
	log.Println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
