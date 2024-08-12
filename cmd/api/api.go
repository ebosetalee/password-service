package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ebosetalee/password-service.git/service/user"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userDB := user.NewRepo(s.db)
	userRoutes := user.NewHandler(userDB)
	userRoutes.RegisterRoutes(subrouter)

	log.Println("Server started on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
