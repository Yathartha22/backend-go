package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"backend-go/backend"
)

func main() {
	backendService := backend.NewService()
	endpoints := backend.MakeEndpoints(backendService)
	ctx := context.Background()

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("backend service is listening on port:", *httpAddr)
		handler := backend.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
