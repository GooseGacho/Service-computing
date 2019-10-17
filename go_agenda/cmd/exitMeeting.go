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

// exitMeetingCmd represents the exitMeeting command
var exitMeetingCmd = &cobra.Command{
	Use:   "exitMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_meeting_, _ := cmd.Flags().GetString("meeting")
		MyExitMeeting(_meeting_)
	},
}

func MyExitMeeting(_meeting_ string) {
	users := entity.READUSERS()
	meetings := entity.READMEETINGS()
	current := entity.GetCurrentUserName()
	if current == "" {
		log.Println("Please log in!")
		return
	}
	//charge user participate meeting
	flag := false //标记用户参加会议
	for i, user := range users {
		if user.Username == current {
			for j, parMeeting := range user.ParticipateMeeting {
				if parMeeting == _meeting_ {
					users[i].ParticipateMeeting = append(user.ParticipateMeeting[:j], user.ParticipateMeeting[j+1:]...)
					flag = true
					break
				}
			}
		}
	}
	//说明用户参加了会议，删除与会人
	if flag {
		for i, meeting := range meetings {
			//不是当前会议
			if meeting.Title != _meeting_ {
				continue
			}
			for j, par := range meeting.Participators {
				//不是与会人
				if par != current {
					continue
				}
				//delete paticipate from meeting log
				meetings[i].Participators = append(meeting.Participators[:j], meeting.Participators[j+1:]...)
				log.Println("Exit " + _meeting_ + " successfully")
				//如果会议没有与会人
				if len(meetings[i].Participators) == 0 {
					//删除会议发起者的会议事件
					var spon = meeting.Sponsor
					for k, user := range users {
						if user.Username == spon {
							for l, sponMeeting := range user.SponsorMeeting {
								//删除发起的会议
								if sponMeeting == _meeting_ {
									users[k].SponsorMeeting = append(user.SponsorMeeting[:l], user.SponsorMeeting[l+1:]...)
								}
							}
						}
					}
					//删除会议
					meetings = append(meetings[:i], meetings[i+1:]...)
					log.Println("Empty meeting! Delete Automitaic!")
				}
			}
		}
		//记录写回
		entity.WRITEUSER(users)
		entity.WRITEMEETINGS(meetings)
		return
	}
	//说明用户没有参加会议
	log.Println("Not Participate Meeting!")
	return
}

func init() {
	rootCmd.AddCommand(exitMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到会议名称[-meeting meeting]
	exitMeetingCmd.Flags().StringP("meeting", "m", "default meeting", "exit meeting participants")
}
