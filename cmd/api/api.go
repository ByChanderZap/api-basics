package api

import (
	"log"
	"net/http"

	"github.com/ByChanderZap/api-basics/cmd/database"
	"github.com/ByChanderZap/api-basics/services/user"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr string
	db   *database.Queries
}

func NewAPIServer(addr string, db *database.Queries) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	//	Getting userRoutes
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(v1Router)

	// Mounting userRoutes v1Router to /api/v1
	router.Mount("/api/v1", v1Router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
