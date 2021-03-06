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

var createType string
var createRewrite bool
var createTioConf bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create code template",
	Long: `Create code template. 
+ GRPC
+ HTTP
+ NSQ.`,
	Run: func(cmd *cobra.Command, args []string) {

		if createTioConf {
			createNewTioConf(createType)
		}

		switch strings.ToLower(createType) {
		case "grpc":
			output("Create GRPC Code Template")
			if err := outputTpl(grpcTpl); err != nil {
				fmt.Printf("Create Implement.go Error. %s\n", err)
				os.Exit(-1)
			}

			if createRewrite {
				os.RemoveAll("rpc")
			}

			if err := os.Mkdir("rpc", 0700); err != nil {
				fmt.Printf("Create rpc dir Error. %s \n", err)
				os.Exit(-1)
			}

		case "http":
			output("Create HTTP Code Template")
		case "nsq":
			output("Create NSQ Code Template")
		default:
			fmt.Println("Soffy, Tio-Cli doesn't support this type code")
			os.Exit(-1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVarP(&createType, "type", "t", "", "Code Template. GRPC/HTTP/NSQ")
	createCmd.PersistentFlags().BoolVarP(&createRewrite, "rewrite", "r", false, "Rewrite implement.go and re-create rpc dir")
	createCmd.PersistentFlags().BoolVarP(&createTioConf, "conf", "c", false, "Rewrite .tio.toml")
	createCmd.MarkPersistentFlagRequired(createType)
}

func createNewTioConf(stype string) {
	var conf = `user=""
[build]
    name=""
    version=""
    api=""
    rate=10
    type="%s"
[deploy]
    url=""`

	file := ".tio.toml"
	os.Remove(file)
	if err := ioutil.WriteFile(file, []byte(fmt.Sprintf(conf, stype)), 0600); err != nil{
		fmt.Println(err)
	}
}

var grpcTpl = `package main

import (
	"context"

	"google.golang.org/grpc"
)

// register 
// Register your grpc server.
func register(s *grpc.Server, srv *server) {
	// Please invoke your grpc register funcion, e.g. rpc.RegisterEchoServer(s, srv)
	
}


// type server struct{}
//
// Server as the truly grpc server instance, it has been declared.
// So please implement GRPC function as the blowing:
// func (s server) Hello(context.Context, *rpc.HelloRequest) (*rpc.HelloResponse, error) {
//		return &rpc.HelloResponse{}, nil
//	}
//
// If you want to initialize server struct, please type code in the flowing function.
//
// func (s server) ServerInit(){
//	 panic("Please Implement Me!")
// }

`

func outputTpl(tpl string) error {
	file := "implement.go"

	if createRewrite {
		os.Remove(file)
	}

	return ioutil.WriteFile(file, []byte(tpl), 0600)
}
