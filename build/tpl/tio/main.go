package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/_ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handler(writer http.ResponseWriter, request *http.Request) {
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	tio_handler(ctx, writer, request)
}
