package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"message": "Hello world"})
	})

	r := http.NewServeMux()
	r.Handle("/hello", handler)
	
	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
	}

	go func(handler http.HandlerFunc) {
		log.Fatal(server.ListenAndServe())
	}(handler)

	go func() {
		<- cs
		defer cancel()
	}()

	<- ctx.Done()
	log.Fatal(server.Shutdown(context.TODO()))
}
