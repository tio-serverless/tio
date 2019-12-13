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
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tio/client/model"
	"tio/client/rpc"
)

var loginCmdURL string
var loginCmdPort int
var loginCmdUser string
var loginCmdPasswd string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login TIO",
	Long: `User should login in before execute every command.
If login success, tio-cli will store useid into $HOME/.tio/tio.toml`,
	Run: func(cmd *cobra.Command, args []string) {

		address, user, passwd, err := getLoginParams()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		output(fmt.Sprintf("url: %s name: %s \n", address, user))

		path := "$HOME/.tio/tio.toml"
		c, err := model.ReadConf(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		uid, err := rpc.Login(address, user, passwd)
		if err != nil {
			fmt.Print("Login Failed! ")
			output(fmt.Sprintf("Error: %s\n", err.Error()))
			fmt.Println()
			os.Exit(-1)
		}

		c.User.Uid = uid

		err = model.UpdateConf(c, path)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&loginCmdURL, "tio", "t", "", "Tio Control URL")
	loginCmd.PersistentFlags().IntVarP(&loginCmdPort, "port", "", 0, "Tio Control URL")
	loginCmd.PersistentFlags().StringVarP(&loginCmdUser, "user", "u", "", "Tio User Name")
	loginCmd.PersistentFlags().StringVarP(&loginCmdPasswd, "passwd", "p", "", "Tio User Passwd")
	loginCmd.MarkPersistentFlagRequired("user")
	loginCmd.MarkPersistentFlagRequired("passwd")
}

func getLoginParams() (address, username, passwd string, err error) {
	var url string
	var port int

	if loginCmdURL == "" {
		if repostry, ok := viper.Get("repostry").(map[string]interface{}); ok {
			if u, ok := repostry["url"]; ok {
				url = u.(string)
			} else {
				err = errors.New("Can not find repostry url. PLease setting it in config file or type it via `-t` ")
				return
			}
		} else {
			err = errors.New("Can not find repostry url. PLease setting it in config file or type it via `-t` ")
			return
		}
	} else {
		url = loginCmdURL
	}

	if loginCmdPort == 0 {
		if repostry, ok := viper.Get("repostry").(map[string]interface{}); ok {
			if p, ok := repostry["port"]; ok {
				port = int(p.(int64))
			} else {
				err = errors.New("Can not find repostry port. PLease setting it in config file or type it via `--port` ")
				return
			}
		} else {
			err = errors.New("Can not find repostry port. PLease setting it in config file or type it via `--port` ")
			return
		}
	} else {
		port = loginCmdPort
	}

	if strings.TrimSpace(loginCmdUser) == "" {
		err = errors.New("User can not be empty! ")
		return
	}

	if strings.TrimSpace(loginCmdPasswd) == "" {
		err = errors.New("Passwd can not be empty! ")
		return
	}

	return fmt.Sprintf("%s:%d", url, port), loginCmdUser, loginCmdPasswd, nil
}
