package get

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sanusomya/birthday-cli/config"
)

func AllBirthdays() []byte {

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

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}
