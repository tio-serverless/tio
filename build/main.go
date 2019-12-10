package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// 1. unzip specify zip file.
// 2. build go source file

var b *B
var (
	zipName string
	control string
	baseImg string
	sid     int
	raw     string
)

func init() {
	flag.StringVar(&zipName, "zip", "", "The Zip File URL")
	flag.StringVar(&control, "control", "", "The Control GRPC Address")
	flag.StringVar(&baseImg, "base", "", "Docker Build Base Image")
	flag.IntVar(&sid, "sid", -1, "The Srv ID")
}

func initBus() {
	b = new(B)
	b.Root = os.Getenv("GOPATH") + "/src"

	if os.Getenv("TIO_DOCKER_USER") != "" {
		b.User = os.Getenv("TIO_DOCKER_USER")
	} else {
		logrus.Fatalln("TIO_DOCKER_USER Empty! ")
	}

	if os.Getenv("TIO_DOCKER_PASSWD") != "" {
		b.Passwd = os.Getenv("TIO_DOCKER_PASSWD")
	} else {
		logrus.Fatalln("TIO_DOCKER_PASSWD Empty! ")
	}

	err := dclientInit()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	info, err := b.DClient.Info()
	if err != nil {
		logrus.Fatalln(err)
	}
	fmt.Println(info.ServerVersion)
}

func main() {
	flag.Parse()
	initBus()

	var err error

	defer func() {
		if err != nil {
			err = faild(control, b.J)
			if err != nil {
				logrus.Errorf("Update status error. %s", err.Error())
			}
			return
		}

		err = succ(control, b.J)
		if err != nil {
			logrus.Errorf("Update status error. %s", err.Error())
		}
		return
	}()

	file := b.Root + "/t.zip"
	logrus.Infof("TIO Build. Zip Path [%s] LocalPath [%s]", zipName, file)

	err = fetch(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = zipAndparser(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	createJob()

	err = build(b.BuildInfo.Name)
	if err != nil {
		logrus.Fatalln(err)
	}

}

func createJob() {
	version := "latest"
	if b.BuildInfo.Version != "" {
		version = b.BuildInfo.Version
	}

	b.J = &job{
		User:  b.UserName,
		Name:  b.BuildInfo.Name,
		Image: fmt.Sprintf("%s:%s-%s", b.Registry, b.BuildInfo.Name, version),
		API:   b.BuildInfo.API,
		Rate:  b.BuildInfo.Rate,
		SType: b.BuildInfo.Stype,
	}
}

func fetch(filepath string) error {
	// Get the data
	resp, err := http.Get(zipName)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
