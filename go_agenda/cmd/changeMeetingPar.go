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

// changeMeetingParCmd represents the changeMeetingPar command
var changeMeetingParCmd = &cobra.Command{
	Use:   "changeMeetingPar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_meeting_, _ := cmd.Flags().GetString("meeting")
		_command_, _ := cmd.Flags().GetString("command")
		_par_, _ := cmd.Flags().GetString("par")
		users := entity.READUSERS()
		meetings := entity.READMEETINGS()
		current := entity.GetCurrentUserName()
		parIndex := -1
		if current == "" {
			log.Println("Please log in!")
			return
		}
		//provide of adding sponsor to meeting
		if current == _par_ {
			log.Println("Cann't add sponsor to meeting as participate")
			return
		}
		//定位需要删除的与会人
		for i, user := range users {
			if user.Username == _par_ {
				parIndex = i
			}
		}
		if parIndex == -1 {
			log.Println("Dont have user named " + _par_)
			return
		}
		for i, meeting := range meetings {
			if meeting.Title == _meeting_ {
				//不是会议发起人，没有权限
				if meeting.Sponsor != current {
					log.Println("Dont have privilege!")
					return
				}
				//删除与会人
				if _command_ == "d" {
					//从会议中删除与会人，找到就删除，没找到则记录错误日志
					for j, partic := range meeting.Participators {
						if partic != _par_ {
							continue
						}
						//从会议中删除与会人
						meetings[i].Participators = append(meeting.Participators[:j], meeting.Participators[j+1:]...)
						//删除该与会人的会议记录
						for k, parMeeting := range users[parIndex].ParticipateMeeting {
							if parMeeting == _meeting_ {
								users[parIndex].ParticipateMeeting = append(users[parIndex].ParticipateMeeting[:k], users[parIndex].ParticipateMeeting[k+1:]...)
							}
						}
						//如果没有与会人
						if len(meetings[i].Participators) == 0 {
							//删除会议发起者的会议事件
							for k, user := range users {
								if user.Username == current {
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
						//记录写回
						log.Println("Delete success!")
						entity.WRITEUSER(users)
						entity.WRITEMEETINGS(meetings)
						return
					}
					log.Println("Dont have particapator name " + _par_)
				} else { //增加与会人
					//与会人查重
					for _, par := range meeting.Participators {
						if par == _par_ {
							log.Println(_par_ + " was participator! Add failed")
							return
						}
					}
					//在会议中加入与会人
					meetings[i].Participators = append(meeting.Participators, _par_)
					//给与会人增加会议事件
					users[parIndex].ParticipateMeeting = append(users[parIndex].ParticipateMeeting, _meeting_)
					log.Println("Add success!")
					//记录写回
					entity.WRITEUSER(users)
					entity.WRITEMEETINGS(meetings)
					return
				}
			}
		}
		log.Println("Dont has this Meeting")
		return
	},
}

func init() {
	rootCmd.AddCommand(changeMeetingParCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeMeetingParCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//得到会议名称[-meeting meeting] 指令名称[-command a/d] 用户名称[-par name]
	changeMeetingParCmd.Flags().StringP("meeting", "m", "default meeting", "change meeting participants")
	changeMeetingParCmd.Flags().StringP("command", "c", "a", "change command participants")
	changeMeetingParCmd.Flags().StringP("par", "p", "default participator", "change par participants")

}
