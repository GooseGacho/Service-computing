package Meeting

import (
	"agenda/entity/User"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Meeting struct {
	Title        string
	Sponsor      string
	Participants []string
	StartTime    time.Time
	EndTime      time.Time
	Id           string
}

//time manage function
func DateToString(date time.Time) string {
	return date.Format("0000-00-00/00:00")
}
func StringToDate(date string) (time.Time, error) {
	the_time, err := time.Parse("0000-00-00/00:00", date)
	return the_time, err
}
func SmallDate(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}
func LargeDate(date1, date2 time.Time) bool {
	return date1.After(date2) || date1.Equal(date2)
}
func TimeContact(newDateS, newDateE, oriDateS, oriDateE time.Time, userName string, meetingId string) bool {
	if LargeDate(newDateS, newDateE) {
		fmt.Println("Invaild time: start time can't be greater than end time")
		return true
	}
	if SmallDate(newDateE, oriDateS) || LargeDate(newDateS, oriDateE) {
		return false
	}
	fmt.Println("time contract with " + userName + "' meeting with id:" + meetingId)
	return true
}

func CreateAMeeting(meeting *Meeting) {
	currentName := User.GetCurUserName()
	if currentName == "" {
		fmt.Println("You haven't logged in")
		return
	}
	var allMeetings map[string]Meeting = GetAllMeetingInfo()
	meeting.Id = strconv.Itoa(len(allMeetings)) //initial id is 0
	meeting.Sponsor = currentName
	meeting.Participants = append(meeting.Participants, currentName)
	allMeetings[meeting.Id] = *meeting

	var allUser map[string]*User.User = GetAllUserInfo()
	//check all participanter exist, and time contract
	for _, pName := range meeting.Participants {
		//check if the user exist
		if _, ok := allUser[pName]; !ok {
			fmt.Println("No such user:" + pName + "!")
			return
		}
		//check if the user's old meeting is contract with the new one
		for _, meetingId := range allUser[pName].ParticipantMeeting {
			if TimeContact(meeting.StartTime, meeting.EndTime, allMeetings[meetingId].StartTime, allMeetings[meetingId].EndTime, pName, meetingId) {
				return
			}
		}
		allUser[pName].ParticipantMeeting = append(allUser[pName].ParticipantMeeting, meeting.Id)
	}
	allUser[currentName].SponsorMeeting = append(allUser[currentName].SponsorMeeting, meeting.Id)

	fout, _ := os.Create("data/Meeting.json")
	defer fout.Close()
	b, _ := json.Marshal(allMeetings)
	fout.Write(b)
	foutuser, _ := os.Create("data/User.json")
	defer foutuser.Close()
	buser, _ := json.Marshal(allUser)
	foutuser.Write(buser)
}

//load all meeting infomation
func GetAllMeetingInfo() map[string]Meeting {
	byteIn, err := ioutil.ReadFile("data/Meeting.json")
	if err != nil {
		log.Fatal(err)
	}
	var allMeetingInfo map[string]Meeting
	json.Unmarshal(byteIn, &allMeetingInfo)
	if allMeetingInfo == nil {
		allMeetingInfo = make(map[string]Meeting)
	}
	return allMeetingInfo
}

//load all user infomation to User.AllUserInfo
func GetAllUserInfo() map[string]*User.User {
	byteIn, err := ioutil.ReadFile("data/User.json")
	if err != nil {
		log.Fatal(err)
	}
	var allUserInfo map[string]*User.User
	json.Unmarshal(byteIn, &allUserInfo)
	return allUserInfo
}
