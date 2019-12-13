/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
	"tio/client/model"
	"tio/client/rpc"
)

var (
	deployCmdDir string
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy code inot Tio",
	Long:  `The dest dir must has a .tio.toml file, which stroe the serverless metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		output(fmt.Sprintf("Deploy all code in %s", deployCmdDir))

		ak, sk, bk, cbu, err := queryToken()
		if err != nil {
			fmt.Printf("Connect With Tio Master Error. %s", err)
			os.Exit(-1)
		}

		dir, name, err := zipDir(deployCmdDir)
		if err != nil {
			fmt.Printf("Zip Error. %s", err)
			os.Exit(-1)
		}

		err = upload(ak, sk, bk, cbu, dir, name)
		if err != nil {
			fmt.Printf("Upload Code Error. %s", err)
			os.Exit(-1)
		}

		fmt.Println("Code Upload Succ. Pleae wait a moment for build and deploy. You can use status command for query progress")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVarP(&deployCmdDir, "dir", "-d", ".", "The Serveless Code Dir")
}

func queryToken() (accessKey, secretKey, bucket, callBackUrl string, err error) {
	if b.TioUrl == "" || b.TioPort == 0 {
		err = errors.New("Please setting repostry metedata in $HOME/.tio.toml")
		return
	}

	if b.UserName == "" {
		err = errors.New("Please login first ")
		return
	}

	return rpc.Token(fmt.Sprintf("%s:%d", b.TioUrl, b.TioPort), b.UserName, b.Passwd)
}

func upload(accessKey, secretKey, bucket, callBackUrl, filePath, fileName string) error {
	localFile := filePath
	key := fileName

	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      callBackUrl,
		CallbackBody:     `{"Message":"{\"key\":\"$(key)\"}"}`,
		CallbackBodyType: "application/json",
	}

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{}

	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	client := http.Client{}

	resumeUploader := storage.NewResumeUploaderEx(&cfg, &storage.Client{Client: &client})
	upToken := putPolicy.UploadToken(mac)

	ret := storage.PutRet{}

	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		//fmt.Println(err)
		return err
	}

	return nil
	//fmt.Println(ret.Key, ret.Hash)
}

func zipDir(path string) (zipDirName, zipFileName string, err error) {
	zipDirName = filepath.Dir(path)

	uid, err := getUserID()
	if err != nil {
		return
	}

	name, err := getMetaData(path)
	if err != nil {
		return
	}

	zipFileName = fmt.Sprintf("%d-%s.zip", uid, name)
	fzip, _ := os.Create(fmt.Sprintf("%s/%s", zipDirName, zipFileName))
	w := zip.NewWriter(fzip)

	files, err := ioutil.ReadDir(zipDirName)
	if err != nil {
		return
	}

	for _, f := range files {
		ff, err := w.Create(f.Name())
		if err != nil {
			return zipDirName, zipFileName, err
		}

		d, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", zipDirName, f.Name()))
		if err != nil {
			return zipDirName, zipFileName, err
		}

		_, err = ff.Write(d)
		if err != nil {
			return zipDirName, zipFileName, err
		}
	}

	w.Close()
	return
}

func getMetaData(path string) (name string, err error) {
	zipDirName := filepath.Dir(path)
	if _, err := os.Stat(fmt.Sprintf("%s/.tio.toml", zipDirName)); os.IsNotExist(err) {
		err = errors.New(fmt.Sprintf("Can not find .tio.toml in %s", zipDirName))
		return name, err
	}

	var m model.MetaData

	_, err = toml.DecodeFile(fmt.Sprintf("%s/.tio.toml", zipDirName), &m)
	if err != nil {
		return
	}

	return m.BuildInfo.Name, nil
}

func getUserID() (id int, err error) {
	path := fmt.Sprintf("%s/.tio/tio.toml", os.Getenv("HOME"))
	c, err := model.ReadConf(path)
	if err != nil {
		return
	}

	return c.User.Uid, nil
}
