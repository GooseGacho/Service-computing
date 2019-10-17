// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	//"fmt"

	"agenda/entity"
	"log"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_userName_, _ := cmd.Flags().GetString("user")
		_password_, _ := cmd.Flags().GetString("pass")
		users := entity.READUSERS()
		for _, user := range users {
			if user.Username == _userName_ {
				//密码匹配
				if user.Password == _password_ {
					entity.SetCurrentUserName(_userName_)
					log.Println("Log in Success!")
					log.Println("Welcome! " + _userName_)
					return
				} else { //密码错误
					log.Println("Warning! Wrong Password")
					return
				}
			}
		}
		log.Println("Warning! Wrong UserName")
		return
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到用户名称[-user userName] 密码[-pass password]
	loginCmd.Flags().StringP("user", "u", "", "log in")
	loginCmd.Flags().StringP("pass", "p", "", "log in")
}
