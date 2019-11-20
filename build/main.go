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
)

func init() {
	b = new(B)
	b.Root = os.Getenv("GOPATH") + "/src"

	err := dclientInit()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	info, err := b.DClient.Info()
	if err != nil {
		logrus.Fatalln(err)
	}
	fmt.Println(info.ServerVersion)

	flag.StringVar(&zipName, "zip", "", "The Zip File URL")
	flag.StringVar(&control, "control", "", "The Control GRPC Address")
	flag.StringVar(&baseImg, "base", "", "Docker Build Base Image")
}

func main() {
	flag.Parse()

	file := b.Root + "/t.zip"
	logrus.Infof("TIO Build. Zip Path [%s] LocalPath [%s]", zipName, file)

	err := fetch(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = zipAndparser(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = build(b.BuildInfo.Name)
	if err != nil {
		logrus.Fatalln(err)
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
