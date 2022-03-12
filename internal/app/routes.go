package app

import (
	"fmt"
	"log"
	"net/http"
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
	listTpl := server.tmpls["list"]

	ans := server.rpcFunc()
	if err := listTpl.Execute(w, ans); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}
