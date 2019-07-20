/*
API server.

See the handlers package for API endpoint documentation.

Binds to ":5000" address.
*/
package main

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Noah-Huppert/slack-stand-up-bot/handlers"

	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
)

// shutdownTimeout is the number of seconds the process is given to gracefully shutdown
const shutdownTimeout int = 120

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	logger := golog.NewStdLogger("slack-stand-up-bot")

	// wg waits until all go routines are finished
	var wg sync.WaitGroup

	// sigs receives interupt signals
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.SIGINT)

	go func() {
		<-sigs
		cancelCtx()
	}()

	// httpRouter holds routes for the API
	httpLogger := logger.GetChild("http")

	httpRouter := mux.NewRouter()
	httpRouter.Handle("/api/v0/health", handlers.HealthHandler{
		Logger: httpLogger.GetChild("health"),
	})
	httpRouter.Handle("/api/v0/slack/webhook", handlers.SlackWebhookHandler{})

	httpServer := http.Server{
		Addr:    ":5000",
		Handler: httpRouter,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := httpServer.ListenAndServe()
		if err != nil {
			logger.Errorf("error serving API: %s", err.Error())
			cancelCtx()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()

		shutdownCtx, _ := context.WithTimeout(context.Background(), shutdownTimeout*time.Seconds)

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			logger.Errorf("error shutting down API: %s", err.Error())
		}
	}()

	// Wait for all go routines to finish
	wg.Wait()
}
