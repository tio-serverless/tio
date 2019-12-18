package main

import (
	"context"
	"log"
	"net"
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

	l, err := net.Listen("tcp4", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, r))

	//srv := &http.Server{
	//
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//
	//log.Fatal(srv.ListenAndServe())
}

func handler(writer http.ResponseWriter, request *http.Request) {
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	tio_handler(ctx, writer, request)
}
