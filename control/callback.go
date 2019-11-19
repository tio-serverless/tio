package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"tio/control/data"
	tio_build_v1 "tio/tgrpc"
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

		url := makePrivateUrl(cui.Key)

		logrus.Debugf("New Request [%s]", cui.Key)

		err = callBuildAgent(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		return
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

func makePrivateUrl(key string) string {

	mac := qbox.NewMac(b.Storage.AcessKey, b.Storage.SecretKey)

	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	return storage.MakePrivateURL(mac, b.Storage.Domain, key, deadline)
}

func callBuildAgent(request string) error {
	conn, err := grpc.Dial(b.BuildAgent, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Build Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_build_v1.NewBuildServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := c.Build(ctx, &tio_build_v1.Request{
		Address: request,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(fmt.Sprintf("Build Agent Return %d", reply.Code))
	}

	return nil
}
