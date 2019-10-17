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

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_password_, _ := cmd.Flags().GetString("pass")
		users := entity.READUSERS()
		current := entity.GetCurrentUserName()
		if current == "" {
			log.Println("Please log in!")
			return
		}
		for i, user := range users {
			if user.Username != current {
				continue
			}
			//密码不正确
			if user.Password != _password_ {
				log.Println("Wrong password!")
				return
			}
			//清空会议
			MyClearMeeting()
			//销户
			users = append(users[:i], users[i+1:]...)
			log.Println("Delete user successfully.")
			//更改当前登陆账户信息
			entity.SetCurrentUserName("")
			break
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到用户密码[-pass password]
	deleteUserCmd.Flags().StringP("pass", "p", "", "delete user")
}
