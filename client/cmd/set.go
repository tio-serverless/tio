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

var setReview bool

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Setting serverless running parameters",
	Long:  `Tio-Cli read .tio.toml and load [deploy] struct data. `,
	Run: func(cmd *cobra.Command, args []string) {
		name := fmt.Sprintf("%d-%s-%s", b.Uid, b.Sname, b.Stype)

		env, err := getNewParam()
		if err != nil {
			fmt.Errorf("Load Env Error. %s", err.Error())
			os.Exit(-1)
		}

		if setReview {
			fmt.Printf(" [%s] New Running Setting : \n", name)
			for key, value := range env {
				fmt.Printf("   %s=%s\n", key, value)
			}
		}

		err = rpc.Update(fmt.Sprintf("%s:%d", b.TioUrl, b.TioPort), name, env)
		if err != nil {
			fmt.Printf("Update Error. %s", err.Error())
			os.Exit(-1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.PersistentFlags().BoolVarP(&setReview, "review", "r", false, "Output new parameters")
}

func getNewParam() (env map[string]string, err error) {
	m, err := getMetaData(os.Getenv("PWD"))
	if err != nil {
		return env, err
	}

	return m.DeployInfo.Env, nil
}
