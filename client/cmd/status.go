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
	"os"

	"github.com/spf13/cobra"
	"tio/client/rpc"
)

var statusName string
var statusLimit int

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Query serverless status",
	Long:  `Tio-cli will load name via reading .tio.toml in the current dir. Of course you can specify serverless name via -s. `,
	Run: func(cmd *cobra.Command, args []string) {
		if statusName == "" {
			var err error
			path := os.Getenv("PWD")
			statusName, err = getServerlessName(path)
			if err != nil {
				fmt.Println("Can not find serverless name in this dir, please use --name specify it. ")
				os.Exit(-1)
			}
		}

		if b.Uid == 0 {
			fmt.Println("Please login first! ")
			os.Exit(-1)
		}

		output(fmt.Sprintf("Query %s status limit return [%d]", statusName, statusLimit))

		ss, err := rpc.Status(fmt.Sprintf("%s:%d", b.TioUrl, b.TioPort), b.Uid, statusLimit, statusName)
		if err != nil {
			fmt.Printf("Query status error! %s \n", err)
			os.Exit(-1)
		}

		fmt.Println()
		fmt.Println("--------------Serverless  Status--------------")
		fmt.Println()
		if len(ss) == 0 {
			fmt.Println(fmt.Sprintf("Can not find [%s] any records.", statusName))
		} else {
			fmt.Println("-------Name----|----Version----|----Status----")
			fmt.Println()
			for _, s := range ss {
				fmt.Printf("%-15s|%-15s|%-15s\n", s.Name, s.Version, s.Status)
			}
		}
		fmt.Println()
		fmt.Println("----------------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	statusCmd.PersistentFlags().StringVarP(&statusName, "name", "n", "", "The Serverless Name")
	statusCmd.PersistentFlags().IntVarP(&statusLimit, "limit", "l", 0, "Limit retrun dataset size")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
