package wish

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/sanusomya/birthday-cli/birthday"
	get "github.com/sanusomya/birthday-cli/cmd/get"
	"github.com/sanusomya/birthday-cli/config"

	"github.com/spf13/cobra"
)

var user []string
var CmdWish = &cobra.Command{
	Use:               "wish",
	SuggestFor:        []string{"w", "wi", "wis"},
	Short:             "use this command to wish the list of all birthdays today",
	Example:           "birthday wish",
	Version:           config.Version,
	Run:               runGet,
	DisableAutoGenTag: true,
}

func runGet(cmd *cobra.Command, args []string) {

	birth := []birthday.Birthday{}
	switch len(user) {
	case 0:
		data := get.AllBirthdaysToday()
		err := json.Unmarshal(data, &birth)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		for _, birthdays := range birth {
			name := birthdays.Name
			sendMessage(name)
		}
	default:
		url := config.GetUrl()

		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(responseData, &birth)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		names := []string{}
		for _, bdays := range birth {
			names = append(names, strings.ToLower(bdays.Name))
		}
		for _, name := range user {
			if slices.Contains(names, strings.ToLower(name)) {
				sendMessage(name)
				continue
			}
			fmt.Printf("this user in not a valid user or not present in database : %v", name)
			fmt.Println()
		}
	}
}

func init() {
	CmdWish.Flags().StringArrayVarP(&user, "user", "u", []string{}, "pass the valid user names to wish them")
}
