package api

import (
	"log"
	"net/http"

	"github.com/ByChanderZap/api-basics/services/cart"
	cartStore "github.com/ByChanderZap/api-basics/services/cart/generated"
	"github.com/ByChanderZap/api-basics/services/product"
	productStore "github.com/ByChanderZap/api-basics/services/product/generated"
	"github.com/ByChanderZap/api-basics/services/user"
	userStore "github.com/ByChanderZap/api-basics/services/user/generated"
	"github.com/ByChanderZap/api-basics/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
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
	uStore := userStore.New(s.db)
	userHandler := user.NewHandler(uStore)
	userHandler.RegisterRoutes(v1Router)

	//	Getting productRoutes
	pStore := productStore.New(s.db)
	productHandler := product.NewHandler(*pStore)
	productHandler.RegisterRoutes(v1Router)

	// Getting cartRoutes
	cartStore := cartStore.New(s.db)
	cartHandler := cart.NewHandler(s.db, cartStore, pStore)
	cartHandler.RegisterRoutes(v1Router)

	//	Mounting userRoutes v1Router to /api/v1
	router.Mount("/api/v1", v1Router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
