package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/seanjh/bingo/backend/api"
)

const (
	defaultDrainTimeout = 15
	defaultTimeout      = 15
	defaultIdleTimout   = 60
)

// via https://github.com/gorilla/mux/tree/v1.8.0#middleware
func main() {
	log.Println("starting server")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*defaultDrainTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := api.NewApiRouter()
	server := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * defaultTimeout,
		ReadTimeout:  time.Second * defaultTimeout,
		IdleTimeout:  time.Second * defaultIdleTimout,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("shutdown error", err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
