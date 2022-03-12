package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"playlist-server/models"

	"github.com/gorilla/mux"
)

//
type Server struct {
	port    string
	tmpls   map[string]*template.Template
	rpcFunc func() models.Playlist
}

// creates a new instance of server.
func New(port string) (*Server, error) {
	server := new(Server)
	server.port = port
	server.tmpls = make(map[string]*template.Template, 1)

	listTmpl, err := template.ParseFiles("./../front/list.gohtml", "./../front/items.gohtml")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates %s", err)
	}

	server.tmpls["list"] = listTmpl
	server.rpcFunc = models.MockRPC

	return server, nil
}

// starts the web server.
func (server *Server) Start(ctx context.Context) error {
	router := mux.NewRouter()
	router.Use(panicMiddleware)
	router.HandleFunc("/", debugRoute).Methods("GET")
	router.HandleFunc("/panic", debugPanic)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.port),
		Handler: router,
	}
	// shutdown on call
	go func() {
		log.Println(srv.Shutdown(ctx))
	}()

	return srv.ListenAndServe()
}
