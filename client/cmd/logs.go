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

var logsBuild, logsRunning, logsFlowing bool
var logsName string

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Query Serverless Log",
	Long:  `Tio-Cli support two logs, build and running. Use -b / -r to switch.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if logsName == "" {
			logsName, err = getServerlessName(os.Getenv("PWD"))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}
		}

		if logsName == "" {
			fmt.Println("Please use -n specify serverless name! ")
			os.Exit(-1)
		}

		name := fmt.Sprintf("%d-%s-%s", b.Uid, b.Sname, b.Stype)
		logs := make(chan string, 1000)

		if logsBuild {
			fmt.Printf("----------[%s-Build-Log]----------\n", logsName)

			if err := rpc.GetBuildLogs(fmt.Sprintf("%s:%d", b.TioUrl, b.TioPort), name, "build", logsFlowing, logs); err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}

			fmt.Println("---------")
			for {
				select {
				case l := <-logs:
					fmt.Println(l)
				}
			}
			os.Exit(0)
		}

		if logsRunning {
			//address, err := queryAgentAddress("deploy")
			//if err != nil {
			//	fmt.Println(err.Error())
			//	os.Exit(-1)
			//}
			//
			//fmt.Println(address)
			fmt.Printf("----------[%s-Running-Log]----------\n", logsName)

		}

	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	logsCmd.PersistentFlags().BoolVarP(&logsBuild, "build", "b", false, "Check Serverless Build Logs")
	logsCmd.PersistentFlags().BoolVarP(&logsRunning, "running", "r", true, "Check Serverless Running Logs")
	logsCmd.PersistentFlags().BoolVarP(&logsFlowing, "flowing", "f", false, "Flowing output the log")
	logsCmd.PersistentFlags().StringVarP(&logsName, "name", "n", "", "Serverless Name")
}

//func queryAgentAddress(stype string) (string, error) {
//	return rpc.GetAgentInfo(fmt.Sprintf("%s:%d", b.TioUrl, b.TioPort), stype)
//}
