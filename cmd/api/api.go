package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ByChanderZap/api-basics/services/product"
	"github.com/ByChanderZap/api-basics/services/user"
	"github.com/ByChanderZap/api-basics/utils"
	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	utils.InitValidator()
	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	//	Getting userRoutes
	userStore := user.New(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1Router)

	//	Getting productRoutes
	productStore := product.New(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(v1Router)

	//	Mounting userRoutes v1Router to /api/v1
	router.Mount("/api/v1", v1Router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
