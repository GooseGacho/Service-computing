package cmd

import (
	"fmt"
	"agenda/entity/User"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		//	fmt.Println("register called by " + username + password + email)
		userInfo := &User.User{username, password, email, make([]string, 0, 5), make([]string, 0, 5)}
		User.RegisterAnUser(userInfo)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "User password")
	registerCmd.Flags().StringP("email", "e", "", "User email")

	}
