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
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var injectType string

// injectCmd represents the inject command
var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject all source code",
	Long:  `If use Inject command, tio-cli will create all source code. You can custom all details.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(injectType) {
		case "grpc":
			if err := ioutil.WriteFile("main.go", []byte(grpcMainTpl), 0600); err != nil {
				fmt.Printf("Create main.go %s", err.Error())
				os.Exit(-1)
			}
		case "http":
			if err := ioutil.WriteFile("main.go", []byte(httpMainTpl), 0600); err != nil {
				fmt.Printf("Create main.go %s", err.Error())
				os.Exit(-1)
			}
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(injectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	injectCmd.PersistentFlags().StringVarP(&injectType, "type", "t", "grpc", "GRPC Serveless")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// injectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var httpMainTpl = `package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Println(r)
		}
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/_ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	l, err := net.Listen("tcp4", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, r))

}

func handler(writer http.ResponseWriter, request *http.Request) {
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	tio_handler(ctx, writer, request)
}

`

var grpcMainTpl = `package main

import (
	"fmt"
	"net"
	"reflect"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hs := health.NewServer()
	// If you see this comment,  Inject command has executed. Grpc service must have a health check, so the bellowing code can not remove.
	hs.SetServingStatus("Tio-GRPC-Service", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, hs)

	reflection.Register(s)
	srv := &server{}

	for i := 0; i < reflect.TypeOf(srv).NumMethod(); i++ {
		method := reflect.TypeOf(srv).Method(i)
		if method.Name == "ServerInit" && method.Type.NumIn() == 1 {
			method.Func.Call([]reflect.Value{
				reflect.ValueOf(srv),
			})
		}
	}

	register(s, srv)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
`
