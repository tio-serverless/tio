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
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

		deployCmdDir = pathconver(deployCmdDir)

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

		os.Remove(fmt.Sprintf("%s/%s", dir, name))
		fmt.Println("Code Upload Succ. Please wait a moment for build and deploy. You can use status command for query progress")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVarP(&deployCmdDir, "dir", "d", ".", "The Serveless Code Dir")
}

func pathconver(p string) string {
	switch p {
	case ".":
		return os.Getenv("PWD")
	}
	return p
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
	localFile := fmt.Sprintf("%s/%s", filePath, fileName)
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
	//zipDirName = filepath.Dir(path)
	zipDirName = path

	uid, err := getUserID()
	if err != nil {
		return
	}

	name, err := getServerlessName(path)
	if err != nil {
		return
	}

	stype, err := getServerlessType(path)
	if err != nil {
		return
	}

	zipFileName = fmt.Sprintf("%d-%s-%s.zip", uid, name, stype)
	err = RecursiveZip(zipDirName, zipFileName)
	if err != nil {
		fmt.Errorf("Zip Error. %s", err.Error())
		return zipDirName, zipFileName, err
	}
	//fzip, _ := os.Create(fmt.Sprintf("%s/%s", zipDirName, zipFileName))
	//w := zip.NewWriter(fzip)
	//
	//files, err := ioutil.ReadDir(zipDirName)
	//if err != nil {
	//	return
	//}
	//
	//for _, f := range files {
	//	ff, err := w.Create(f.Name())
	//	if err != nil {
	//		return zipDirName, zipFileName, err
	//	}
	//
	//	d, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", zipDirName, f.Name()))
	//	if err != nil {
	//		return zipDirName, zipFileName, err
	//	}
	//
	//	_, err = ff.Write(d)
	//	if err != nil {
	//		return zipDirName, zipFileName, err
	//	}
	//}
	//
	//w.Close()
	return
}

func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = myZip.Close()
	if err != nil {
		return err
	}
	return nil
}

func getMetaData(path string) (m model.MetaData, err error) {
	m = model.MetaData{}

	if _, err := os.Stat(fmt.Sprintf("%s/.tio.toml", path)); os.IsNotExist(err) {
		err = errors.New(fmt.Sprintf("Can not find .tio.toml in %s", path))
		return m, err
	}

	_, err = toml.DecodeFile(fmt.Sprintf("%s/.tio.toml", path), &m)
	if err != nil {
		return
	}

	return m, nil
}

func getServerlessName(path string) (name string, err error) {
	m, err := getMetaData(path)
	if err != nil {
		return "", err
	}

	return m.BuildInfo.Name, nil
}

func getServerlessType(path string) (stype string, err error) {
	m, err := getMetaData(path)
	if err != nil {
		return stype, err
	}

	return m.BuildInfo.Stype, nil
}
func getServrelessVersion(path string) (version string, err error) {
	m, err := getMetaData(path)
	if err != nil {
		return "", err
	}

	return m.BuildInfo.Version, nil
}

func getUserID() (id int, err error) {
	return b.Uid, nil
}
