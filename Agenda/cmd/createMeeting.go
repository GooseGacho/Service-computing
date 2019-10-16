package cmd

import (
	"agenda/entity/Meeting"
	"fmt"
	"github.com/spf13/cobra"
)

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetStringSlice("participants")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		timeS, _ := Meeting.StringToDate(startTime)
		timeE, _ := Meeting.StringToDate(endTime)
		Meeting.CreateAMeeting(&Meeting.Meeting{title, "", participants, timeS, timeE, ""})
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "title")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "startTime")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "User endTime")
}