package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RunHTTPServerOnAddr(port string, addRoutes func(r chi.Router) chi.Router) {
	rootRouter := chi.NewRouter()
	addMiddleware(rootRouter)

	apiRouter := chi.NewRouter()
	rootRouter.Mount("/api", addRoutes(apiRouter))

	log.Println(fmt.Sprintf("Starting server on port: %s", port))
	log.Fatal(http.ListenAndServe(port, rootRouter))
}

func addMiddleware(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	addCorsMiddleware(router)
	addAuthMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	// Add the CORS middleware
}

func addAuthMiddleware(router *chi.Mux) {
	// Add the Authentication middleware
}
