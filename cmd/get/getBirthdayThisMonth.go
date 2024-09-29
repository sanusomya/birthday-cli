package add

import (
	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func allBirthdaysMonth() {

	url := config.GetUrl()
	now := time.Now()
	month := now.Month().String()
	month = strings.ToLower(month[:3])
	url = url + "month?month=" + month

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	utils.Print(responseData)
}
