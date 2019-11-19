package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

// 1. unzip specify zip file.
// 2. build go source file

var b *B
var (
	zipName string
)

func init() {
	b = new(B)
	b.Root = os.Getenv("GOPATH") + "/src"

	flag.StringVar(&zipName, "zip", "", "The Zip File Name")

}

func main() {
	flag.Parse()

	logrus.Infof("TIO Build. Zip Path [%s]", zipName)

	err := zipAndparser(zipName)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = build(b.Name)
	if err != nil {
		logrus.Fatalln(err)
	}
}
