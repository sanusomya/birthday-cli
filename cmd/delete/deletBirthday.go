package delete

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

var CmdDelete = &cobra.Command{
	Use:               "delete",
	Aliases:           []string{"del"},
	SuggestFor:        []string{"del", "d", "delet"},
	Short:             "use this command to delete the entry from birthdays",
	Example:           "birthday delete -name <data> -day <data> -month <data> -mobile <data>",
	Version:           config.Version,
	PreRun:            utils.ValidFlags,
	Run:               runDelete,
	DisableAutoGenTag: true,
}

func runDelete(cmd *cobra.Command, args []string) {


	dateAsInt, _ := strconv.Atoi(date)
	mobileAsInt, _ := strconv.Atoi(phone)

	url := config.GetUrl()

	temp := birthday.Birthday{}
	temp.Name = name
	temp.Month = month
	temp.Date = int8(dateAsInt)
	temp.Mobile = int64(mobileAsInt)

	jsonValue, _ := json.Marshal(temp)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(jsonValue))
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
	CmdDelete.Flags().StringVarP(&name, "name", "n", "", "name of person")
	CmdDelete.Flags().StringVarP(&month, "month", "m", "", "month")
	CmdDelete.Flags().StringVarP(&date, "date", "d", "", "date")
	CmdDelete.Flags().StringVarP(&phone, "phone", "p", "", "phone number")
}
