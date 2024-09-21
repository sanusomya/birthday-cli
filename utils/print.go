package utils

import (
	"birthday-cli/birthday"
	"encoding/json"
	"fmt"
)

func PrintSingleObject(res []byte) {
	bday := birthday.Birthday{}
	err := json.Unmarshal(res, &bday)
	if err != nil {
		fmt.Println(string(res))
	}

	fmt.Printf("Name: %v, Date: %v, Month: %v", bday.Name, bday.Date, bday.Month)
	fmt.Println()
}

func Print(res []byte) {
	birthday := []birthday.Birthday{}
	err := json.Unmarshal(res, &birthday)
	if err != nil {
		fmt.Println(string(res))
	}
	for _, bday := range birthday {
		fmt.Printf("Name: %v, Date: %v, Month: %v", bday.Name, bday.Date, bday.Month)
		fmt.Println()
	}

}
