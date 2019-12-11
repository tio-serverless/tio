package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"tio/control/data"
	"tio/control/db"
	"tio/database/model"
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
			logrus.Error(err.Error())
			return
		}

		err = json.Unmarshal(content, &cu)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		cui := data.CodeUploadInfo{}

		err = json.Unmarshal([]byte(cu.Message), &cui)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		url := makePrivateUrl(cui.Key)

		logrus.Debugf("New Request [%s]", cui.Key)

		err = callBuildAgent(cui.Key, url)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		return
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         add,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func makePrivateUrl(key string) string {

	mac := qbox.NewMac(b.Storage.AcessKey, b.Storage.SecretKey)

	deadline := time.Now().Add(time.Second * 3600).Unix()
	return storage.MakePrivateURL(mac, b.Storage.Domain, key, deadline)
}

func callBuildAgent(key, request string) error {
	var err error

	uid, name := splitUidAndSrvName(key)
	if uid == 0 {
		return errors.New(fmt.Sprintf("Can not split uid and srv name from [%s]. ", key))
	}

	sid, err := db.SaveNewSrv(b, uid, name)
	if err != nil {
		return errors.New(fmt.Sprintf("Save Srv Error [%s]. ", err))
	}

	defer func() {
		if err != nil {
			db.UpdateSrvStatus(b, sid, model.SrvBuildFailed)
		}
	}()

	conn, err := grpc.Dial(b.BuildAgent, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect Build Service Error: %s", err.Error()))
	}

	defer conn.Close()

	c := tio_build_v1.NewBuildServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reply, err := c.Build(ctx, &tio_build_v1.Request{
		Name:    strings.Split(key, ".")[0],
		Address: request,
		Sid:     int32(sid),
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(fmt.Sprintf("Build Agent Return %s", reply.Msg))
	}

	return nil
}

func splitUidAndSrvName(fileName string) (int, string) {
	var uid int
	var name string
	fs := strings.Split(fileName, "-")
	if len(fs) != 2 {
		return uid, name
	}

	uid, err := strconv.Atoi(fs[0])
	if err != nil {
		return uid, name
	}

	name = fs[1]

	return uid, name
}
