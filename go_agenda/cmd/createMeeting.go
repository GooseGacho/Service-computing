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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("meeting")
		month, _ := cmd.Flags().GetInt("month")
		day, _ := cmd.Flags().GetInt("day")
		time, _ := cmd.Flags().GetInt("time")
		com, _ := cmd.Flags().GetString("command")
		users := entity.READUSERS()
		meetings := entity.READMEETINGS()
		current := entity.GetCurrentUserName()
		if current == "" {
			log.Println("Please log in!")
			return
		}
		//添加新的时间
		if com == "update" {
			for i, meeting := range meetings {
				if meeting.Title != title {
					continue
				}
				if meeting.Sponsor != current {
					log.Println("Wrong! You aren't the Sponsor!")
					return
				}
				for j, readyTime := range meeting.MeetingTime {
					//增加当天不同的时间段
					if readyTime.Day == day && readyTime.Month == month {
						for _, tid := range readyTime.TimeID {
							//时间重复
							if tid == time {
								log.Println("Wrong! You had apply this time")
								return
							}
						}
						meetings[i].MeetingTime[j].TimeID = append(readyTime.TimeID, time)
						log.Println("Apply Success!")
						//记录写回
						entity.WRITEUSER(users)
						entity.WRITEMEETINGS(meetings)
						return
					} else { //增加不同的时间
						newTime := entity.Time{TimeID: []int{0: time}, Day: day, Month: month}
						meetings[i].MeetingTime = append(meeting.MeetingTime, newTime)
						log.Println("Apply Success!")
						//记录写回
						entity.WRITEUSER(users)
						entity.WRITEMEETINGS(meetings)
						return
					}
				}
			}
			log.Println("Wrong! You should use -command new")
		} else { //新建会议事项
			//名称、时间查重
			for _, meeting := range meetings {
				//名称查重
				if meeting.Title == title {
					log.Println("Wrong! You should use -command update")
					return
				}
				//时间查重
				for _, readyTime := range meeting.MeetingTime {
					//同一天不同时间段查重
					if readyTime.Day == day && readyTime.Month == month {
						for _, tid := range readyTime.TimeID {
							//时间重复
							if tid == time {
								log.Println("Wrong! Time had applied!")
								return
							}
						}
					}
				}
			}
			//创建新的会议事件并加入会议列表
			newTime := entity.Time{TimeID: []int{time}, Day: day, Month: month}
			newMeeting := entity.Meeting{Title: title, Sponsor: current, Participators: []string{}, MeetingTime: []entity.Time{newTime}}
			meetings = append(meetings, newMeeting)
			//为操作者添加会议事件
			for i, user := range users {
				if user.Username == current {
					users[i].SponsorMeeting = append(users[i].SponsorMeeting, title)
					break
				}
			}
			log.Println("Apply Success! Please add Paticipators!")
			//记录写回
			entity.WRITEUSER(users)
			entity.WRITEMEETINGS(meetings)
			// log.Println(len(meetings))
			// for _, user := range users {
			// 	log.Println(user.Username)
			// 	log.Println(user.SponsorMeeting)
			// }
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//得到会议名称[-meeting meeting] [-Month month] [-Day day] [-time time]1/2/3/4 [-command command]update/new 增加新的时间/创建新的会议
	createMeetingCmd.Flags().StringP("meeting", "m", "default meeting", "create meeting participants")
	createMeetingCmd.Flags().IntP("month", "M", 1, "create month participants")
	createMeetingCmd.Flags().IntP("day", "d", 1, "create day participants")
	createMeetingCmd.Flags().IntP("time", "t", 1, "create time participants")
	createMeetingCmd.Flags().StringP("command", "c", "a", "create command participants")
}
