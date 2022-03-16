package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func debugRoute(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "debug route"); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Printf("debug handler failed: %s", err)
	}
}

func debugPanic(w http.ResponseWriter, r *http.Request) {
	panic("panic route")
}

func (server *Server) playlist(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		listTpl := server.tmpls["list"]

		ans, err := server.rpcFunc(ctx, vars["playlistId"])
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err := listTpl.Execute(w, ans); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/index.html")
}
