package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rayhanm17/ecom-api-go/routes/user"
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
	r := chi.NewRouter()  // create a new router
	r.Mount("/api/v1", r) // Path Prefix

	userHandler := user.NewHandler()
	userHandler.UserRoutes(r)

	return http.ListenAndServe(s.addr, r)
}
