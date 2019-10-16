package User

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Username           string
	Password           string
	Email              string
	SponsorMeeting     []string
	ParticipantMeeting []string
}

//register an  user with name, password, email
func RegisterAnUser(user *User) {
	AllUserInfo := GetAllUserInfo()
	flog, err := os.OpenFile("data/input_output.log", os.O_APPEND|os.O_WRONLY, 0600)
	defer flog.Close()
	check(err)
	logger := log.New(flog, "", log.LstdFlags)
	logger.Printf("agenda register -u %s -p %s -e %s", user.Username, user.Password, user.Email)

	if _, ok := AllUserInfo[user.Username]; !ok {
		AllUserInfo[user.Username] = *user
		os.Stdout.WriteString("register succeed!\n")
		logger.Print("register succeed!\n")
	} else {
		os.Stdout.WriteString("The userName have been registered\n")
		logger.Print("The userName have been registered\n")
	}
}

//Get all user infomation
func GetAllUserInfo() map[string]User {

	byteIn, err := ioutil.ReadFile("data/User.json")
	check(err)
	var allUserInfo map[string]User
	json.Unmarshal(byteIn, &allUserInfo)
	return allUserInfo
}

func check(r error) {
	if r != nil {
		log.Fatal(r)
	}
}
