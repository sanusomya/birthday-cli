package add

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"

	"github.com/spf13/cobra"
)

var today bool
var month bool
var CmdGet = &cobra.Command{
	Use:               "get",
	SuggestFor:        []string{"g", "gt", "et"},
	Short:             "use this command to get the list of all birthdays",
	Example:           "birthday get",
	Version:           config.Version,
	Run:               runGet,
	DisableAutoGenTag: true,
}

func runGet(cmd *cobra.Command, args []string) {

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

	if today {
		AllBirthdaysToday()
		return
	}
	if month {
		allBirthdaysMonth()
		return
	}
	utils.Print(responseData)
}

func init() {
	CmdGet.Flags().BoolVarP(&today, "today", "t", false, "today")
	CmdGet.Flags().BoolVarP(&month, "month", "m", false, "month")
}
