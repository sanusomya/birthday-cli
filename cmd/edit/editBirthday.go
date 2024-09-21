package edit

import (
	"birthday-cli/birthday"
	"birthday-cli/config"
	"birthday-cli/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var name string
var month string
var date string
var phone string

var CmdEdit = &cobra.Command{
	Use:               "edit",
	SuggestFor:        []string{"edt", "ed"},
	Short:             "use this command to delete the entry from birthdays",
	Example:           "birthday edit -name <data> -day <data> -month <data> -phone <data>",
	Version:           config.Version,
	PreRun:            utils.ValidFlags,
	Run:               runEdit,
	DisableAutoGenTag: true,
}

func runEdit(cmd *cobra.Command, args []string) {
	dateAsInt, _ := strconv.Atoi(date)
	mobileAsInt, _ := strconv.Atoi(phone)

	url := config.GetUrl()

	temp := birthday.Birthday{}
	temp.Name = name
	temp.Month = month
	temp.Date = int8(dateAsInt)
	temp.Mobile = int64(mobileAsInt)

	jsonValue, _ := json.Marshal(temp)

	url = url + "edit/?name=" + name + "&mobile=" + phone

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonValue))
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
	utils.PrintSingleObject(responseData)

}

func init() {
	CmdEdit.PersistentFlags().StringVarP(&name, "name", "n", "", "name of person")
	CmdEdit.PersistentFlags().StringVarP(&month, "month", "m", "", "month")
	CmdEdit.PersistentFlags().StringVarP(&date, "date", "d", "", "date")
	CmdEdit.PersistentFlags().StringVarP(&phone, "phone", "p", "", "phone number")
	CmdEdit.AddCommand(cmdEditName)
	CmdEdit.AddCommand(cmdEditMobile)
}
