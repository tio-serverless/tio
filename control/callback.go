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
			logrus.Errorf("Parse Json Error: %s, Body: %s", err.Error(), string(content))
			return
		}

		cui := data.CodeUploadInfo{}

		err = json.Unmarshal([]byte(cu.Message), &cui)
		if err != nil {
			logrus.Errorf("Parse Upload Info Error: %s, CodeUpload: %v", err.Error(), cu)
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

// callBuildAgent
// 从七牛云接受上传回调事件,然后记录DB后调用构建事件
func callBuildAgent(key, request string) error {
	var err error

	uid, name, stype := splitUidAndSrvName(key)
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
		Name:      trimTimestamp(key),
		Address:   request,
		Sid:       int32(sid),
		BuildType: stype,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(fmt.Sprintf("Build Agent Return %s", reply.Msg))
	}

	return db.UpdateSrvStatus(b, sid, model.SrvBuilding)
}

// splitUidAndSrvName 从文件名中获取用户ID、服务名称和服务类型
// 文件名按照  id-name-type-timestamp.zip 规则拼装
func splitUidAndSrvName(fileName string) (int, string, string) {
	var uid int
	var name string
	var stype string

	if !strings.HasSuffix(fileName, ".zip") {
		return uid, name, stype
	}

	fileName = strings.Split(fileName, ".")[0]

	fs := strings.Split(fileName, "-")
	if len(fs) != 4 {
		return uid, name, stype
	}

	uid, err := strconv.Atoi(fs[0])
	if err != nil {
		return uid, name, stype
	}

	name = fs[1]
	stype = fs[2]
	return uid, name, stype
}

func trimTimestamp(filename string) string {
	preName := strings.Split(filename, ".")[0]
	buildName := strings.Split(preName, "-")

	return strings.Join(buildName[:len(buildName)-1], "-")
}
