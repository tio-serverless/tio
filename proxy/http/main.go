package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var si *svcImplement

func main() {
	var err error
	si, err = newSI()
	if err != nil {
		logrus.Fatal(err)
	}

	go start(si, 8000)
	setupRoute(si)
}

func setupRoute(loader dataLoader) {
	router := mux.NewRouter()
	router.HandleFunc("/{url}", func(w http.ResponseWriter, r *http.Request) {
		Proxy(loader, w, r)
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
