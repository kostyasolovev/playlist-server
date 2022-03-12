package app

import (
	"fmt"
	"log"
	"net/http"
)

func debugRoute(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "debug route"); err != nil {
		log.Printf("debug handler failed: %s", err)
	}
}

func debugPanic(w http.ResponseWriter, r *http.Request) {
	panic("panic route")
}
