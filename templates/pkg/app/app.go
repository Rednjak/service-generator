package app

import (
	"log"
	"net/http"

	"github.com/Rednjak/zlat/pkg/repository"
	"github.com/Rednjak/zlat/pkg/service"

	"github.com/go-chi/chi"
)

type server struct {
	router  *chi.Mux
	service service.Service
}

// Run will initialize all the dependencies and return a server struct which can be used to start the server
func Run() server {
	s := server{}
	storage := repository.SetupDatabase()
	service := service.SetupService(storage)

	s.router = chi.NewRouter()
	s.service = service

	s.routes()
	return s
}

// StartServer starts the server
func (s *server) StartServer() {
	log.Println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
