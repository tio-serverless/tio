package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"tio/control/data"
)

func restWeb() {
	add := fmt.Sprintf("0.0.0.0:%d", b.RestPort)
	logrus.Infof("Rest Web Listen %s", add)

	router := mux.NewRouter()

	router.HandleFunc("/code/upload", func(w http.ResponseWriter, r *http.Request) {
		cu := data.CodeUpload{}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = json.Unmarshal(content, &cu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cui := data.CodeUploadInfo{}

		err = json.Unmarshal([]byte(cu.Message), &cui)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

	})

	srv := &http.Server{
		Handler: router,
		Addr:    add,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
