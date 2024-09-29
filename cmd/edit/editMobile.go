package edit

import (
	"github.com/sanusomya/birthday-cli/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var cmdEditMobile = &cobra.Command{
	Use:        "mobile",
	Aliases:    []string{"mob"},
	SuggestFor: []string{"mobi", "mo", "moble"},
	Short:      "use this command to edit name of the entry from birthdays",
	Example:    "birthday edit -name <data> -day <data> -month <data> -mobile <data> editname <data>",
	Version:    config.Version,
	//PreRun:            utils.ValidFlags,
	Run:               runEditMobile,
	DisableAutoGenTag: true,
}

func runEditMobile(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		fmt.Println("you need an argument first. try -h")
		os.Exit(1)
	}
	name, _ := cmd.Flags().GetString("name")
	mobile, _ := cmd.Flags().GetString("phone")

	url := config.GetUrl()

	client := &http.Client{}
	url = url + "edit/number?name=" + name + "&mobile=" + mobile
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(args[0])))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
