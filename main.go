package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/VishalHasija/Product-API.git/handlers"
)

func main() {

	l := log.New(os.Stdout, "Product-API", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewBye(l)
	mux := http.NewServeMux()
	mux.Handle("/goodbye", gh)
	mux.Handle("/", hh)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	l.Println("Received terminate, graceful shutdown ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
