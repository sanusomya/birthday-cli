package add

import (
	"github.com/sanusomya/birthday-cli/birthday"
	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"
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

var CmdAdd = &cobra.Command{
	Use:               "add",
	SuggestFor:        []string{"ad", "d", "a"},
	Short:             "use this command to add to the list of all birthdays",
	Example:           "birthday add -name <data> -day <data> -month <data> -phone <data>",
	Version:           config.Version,
	PreRun:            utils.ValidFlags,
	Run:               runAdd,
	DisableAutoGenTag: true,
}

func runAdd(cmd *cobra.Command, args []string) {


	dateAsInt, _ := strconv.Atoi(date)
	mobileAsInt, _ := strconv.Atoi(phone)

	url := config.GetUrl()

	temp := birthday.Birthday{}
	temp.Name = name
	temp.Month = month
	temp.Date = int8(dateAsInt)
	temp.Mobile = int64(mobileAsInt)

	jsonValue, _ := json.Marshal(temp)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

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
	CmdAdd.Flags().StringVarP(&name, "name", "n", "", "name of person")
	CmdAdd.Flags().StringVarP(&month, "month", "m", "", "month")
	CmdAdd.Flags().StringVarP(&date, "date", "d", "", "date")
	CmdAdd.Flags().StringVarP(&phone, "phone", "p", "", "phone number")
}
