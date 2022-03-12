package app

import (
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

func (server *Server) playlist(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	listTpl := server.tmpls["list"]

	ans := server.rpcFunc(vars["playlistId"])
	if err := listTpl.Execute(w, ans); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
	// fmt.Fprintf(w, "Category: %v\n")
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./../front/index.html")
}
