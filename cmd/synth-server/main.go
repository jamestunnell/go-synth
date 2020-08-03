package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/mux"
	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/pkg/api"
)

func main() {

	usage := `Audio synthesis server.

Usage:
	synth-server [options] <port>

Options:
	-h, --help`

	args, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%# v", pretty.Formatter(args))

	port, err := strconv.Atoi(args["<port>"].(string))
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	api.AddRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	runServer(server)
}

func runServer(server *http.Server) {
	interrupt := make(chan os.Signal, 1)
	defer close(interrupt)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go serverListen(server)

	// block until signal is received on interrupt channel
	sig := <-interrupt

	log.Printf("interrupted by system signal: %s", sig.String())

	// attempt a graceful shutdown using context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Print("shut down server")
}

func serverListen(server *http.Server) {
	log.Printf("starting server on %s", server.Addr)

	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Printf("server error: %v", err)
	}

	log.Print("server stopped")
}
