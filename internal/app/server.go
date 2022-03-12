package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

// creates a new instance of server.
func New(port string) (*Server, error) {
	return &Server{
		port: port,
	}, nil
}

// starts the web server.
func (server *Server) Start(ctx context.Context) error {
	router := mux.NewRouter()
	router.Use(panicMiddleware)
	router.HandleFunc("/", debugRoute)
	router.HandleFunc("/panic", debugPanic)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.port),
		Handler: router,
	}
	// shutdown on call
	go func() {
		<-ctx.Done()
		log.Println(srv.Shutdown(ctx))
	}()

	return srv.ListenAndServe()
}
