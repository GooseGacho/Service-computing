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
	"agenda/entity"
	"log"

	"github.com/spf13/cobra"
)

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "deleteMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_meeting_, _ := cmd.Flags().GetString("meeting")
		current := entity.GetCurrentUserName()
		if current == "" {
			log.Println("Please log in!")
			return
		}
		MyDeleteMeeting(_meeting_)
	},
}

func MyDeleteMeeting(_meeting_ string) {
	users := entity.READUSERS()
	meetings := entity.READMEETINGS()
	current := entity.GetCurrentUserName()
	for i, meeting := range meetings {
		if meeting.Title == _meeting_ {
			//判断是否是会议发起人
			if meeting.Sponsor != current {
				log.Println("Dont have privilage!")
				return
			}
			//删除所有与会人及发起者的会议记录
			for _, par := range meeting.Participators {
				for j, user := range users {
					//删除该与会人的会议记录
					if user.Username == par {
						for l, parMeeting := range user.ParticipateMeeting {
							if parMeeting == _meeting_ {
								users[j].ParticipateMeeting = append(user.ParticipateMeeting[:l], user.ParticipateMeeting[l+1:]...)
							}
						}
					}
				}
			}
			//delete Sponsor meeting log
			for j, user := range users {
				//delete Sponsor meeting log
				if user.Username == current {
					for l, sponMeeting := range user.SponsorMeeting {
						if sponMeeting == _meeting_ {
							users[j].SponsorMeeting = append(user.SponsorMeeting[:l], user.SponsorMeeting[l+1:]...)
						}
					}
				}
			}
			//删除会议
			meetings = append(meetings[:i], meetings[i+1:]...)
			log.Println("Delete " + _meeting_ + " Success!")
			//记录写回
			entity.WRITEUSER(users)
			entity.WRITEMEETINGS(meetings)
			return
		}
	}
	//如果遍历结束都没有返回，证明会议不存在，错误写回日志
	log.Println("Dont have this Meeting")
	return
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到会议名称[-meeting meeting]
	deleteMeetingCmd.Flags().StringP("meeting", "m", "default meeting", "delete meeting participants")
}
