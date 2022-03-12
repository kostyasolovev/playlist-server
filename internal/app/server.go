package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

func New(port string) (*Server, error) {
	return &Server{
		port: port,
	}, nil
}

func (server *Server) Start(ctx context.Context) error {
	router := mux.NewRouter()
	router.Use(panicMiddleware)
	router.HandleFunc("/", debugRoute)
	router.HandleFunc("/panic", debugPanic)

	return http.ListenAndServe(fmt.Sprintf(":%s", server.port), router)
}
