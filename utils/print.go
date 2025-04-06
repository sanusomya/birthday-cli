package utils

import (
	"github.com/sanusomya/birthday-cli/birthday"
	"encoding/json"
	"fmt"
)

func PrintSingleObject(res []byte) {
	bday := birthday.Birthday{}
	err := json.Unmarshal(res, &bday)
	if err != nil {
		fmt.Println(string(res))
		return
	}
	if bday.Cell == 0{
		fmt.Println(string(res))
		return
	}
	fmt.Printf("Name: %v, Date: %v, Month: %v", bday.Person, bday.Birthdate, bday.Birthmonth)
	fmt.Println()
}

func Print(res []byte) {
	birthday := []birthday.Birthday{}
	err := json.Unmarshal(res, &birthday)
	if err != nil {
		fmt.Println(string(res))
	}
	for _, bday := range birthday {
		fmt.Printf("Name: %v, Date: %v, Month: %v", bday.Person, bday.Birthdate, bday.Birthmonth)
		fmt.Println()
	}

}
