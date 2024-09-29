package add

import (
	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func allBirthdaysToday() {

	url := config.GetUrl()
	now := time.Now()
	date := strconv.Itoa(now.Day())
	month := now.Month().String()
	month = strings.ToLower(month[:3])
	url = url + "today?month=" + month + "&date=" + date

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
