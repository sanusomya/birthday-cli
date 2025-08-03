package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sanusomya/birthday-cli/birthday"
	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"

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
	temp.Person = name
	temp.Birthmonth = month
	temp.Birthdate = int8(dateAsInt)
	temp.Cell = int64(mobileAsInt)

	jsonValue, _ := json.Marshal(temp)

	headers := map[string]string{
		"Authorization": fmt.Sprintf("%s", os.Getenv("token")),
		"Content-Type":  "application/json",
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
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
