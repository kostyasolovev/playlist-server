package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"playlist-server/internal/app"
	"syscall"
)

func main() {
	defaultPort := "8080"
	newPort := flag.String("p", defaultPort, "port to listen")
	flag.Parse()

	ctx, finish := context.WithCancel(context.TODO())

	// graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-exitChan
		finish()
	}()
	// init server
	server, err := app.New((*newPort))
	if err != nil {
		log.Fatalf("app.New failed %s", err)
	}
	// run server
	if err := server.Start(ctx); err != nil {
		log.Println("server error", err)
	}
}
