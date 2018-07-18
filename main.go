package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anubhavmishra/hello-oscon/handlers"
	"github.com/braintree/manners"
)

const version = "1.0.0"

func main() {

	var httpBindAddr = "0.0.0.0"
	var httpPort = os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	httpAddr := fmt.Sprintf("%s:%s", httpBindAddr, httpPort)
	log.Println("Starting hello-oscon app.....")

	mux := http.NewServeMux()
	mux.Handle("/", handlers.HelloWorldHandler("hello oscon, I am mishra.", version))
	mux.HandleFunc("/healthz", handlers.HealthCheck)

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = handlers.LoggingHandler(mux)

	errChan := make(chan error, 10)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
