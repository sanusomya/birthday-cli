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

	headers := map[string]string{
		"Authorization": fmt.Sprintf("%s", os.Getenv("token")),
		"Content-Type":  "application/json",
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

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
