package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/kostyasolovev/youtube_pb_api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"playlist-server/models"
)

//
type Server struct {
	port     string
	tmpls    map[string]*template.Template
	rpcFunc  func(context.Context, string) (models.Playlist, error)
	grpcConn *grpc.ClientConn
	grpcCli  api.YoutubePlaylistClient
}

// creates a new instance of server.
func New(port string) (*Server, error) {
	server := new(Server)
	server.port = port
	server.tmpls = make(map[string]*template.Template, 1)

	listTmpl, err := template.ParseFiles("./front/list.gohtml", "./front/items.gohtml")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates %w", err)
	}

	server.tmpls["list"] = listTmpl

	server.rpcFunc = server.registerDefaultListGRPCFunc()

	return server, nil
}

// starts the web server.
func (server *Server) Start(ctx context.Context) error {
	if err := server.dialGRPC(ctx, "8081"); err != nil {
		return err
	}

	if err := server.registerGRPCClient(ctx); err != nil {
		return err
	}

	router := mux.NewRouter()
	router.Use(panicMiddleware)
	router.HandleFunc("/debug", debugRoute).Methods("GET")
	router.HandleFunc("/panic", debugPanic)
	router.HandleFunc("/", mainRoute)
	router.HandleFunc("/{playlistId}/list", server.playlist(ctx))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.port),
		Handler: router,
	}
	// shutdown on call
	go func() {
		<-ctx.Done()
		log.Println(srv.Shutdown(ctx))
	}()

	return errors.Wrap(srv.ListenAndServe(), "server failed")
}
