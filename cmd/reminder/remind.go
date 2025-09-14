package reminder

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sanusomya/birthday-cli/birthday"
	get "github.com/sanusomya/birthday-cli/cmd/get"
	"github.com/sanusomya/birthday-cli/config"
	"github.com/sanusomya/birthday-cli/utils"

	"github.com/spf13/cobra"
)

var dayRange int

var CmdRmd = &cobra.Command{
	Use:               "remind",
	SuggestFor:        []string{"r", "rmd", "rimnd", "remindr"},
	Short:             "use this command to dend a remainder for all birthdays today",
	Example:           "birthday wish",
	Version:           config.Version,
	Run:               runGet,
	DisableAutoGenTag: true,
}

func runGet(cmd *cobra.Command, args []string) {

	birth := []birthday.Birthday{}
	data := get.AllBirthdays()
	err := json.Unmarshal(data, &birth)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	for _, birthdays := range birth {
		now := time.Now()
		month := int(now.Month())
		day := now.Day()
		bdayMon := utils.MonthToNumber(birthdays.Birthmonth)
		if dayRange > 6 || dayRange < 0 {
			fmt.Println("Range cannot be more than 6 or a negative entity. Please enter a valid range.")
			return
		}
		if month == bdayMon {
			if (int(birthdays.Birthdate)-day) <= dayRange && (int(birthdays.Birthdate)-day) >= 0 {
				sendHtmlMail(birthdays.Person, os.Getenv("mail"), birthdays.Birthdate, birthdays.Birthmonth, birthdays.Cell)
			}
		}
	}
}

func init() {
	CmdRmd.Flags().IntVar(&dayRange, "range", 0, "use this flag to set a range of days to send reminders for, \neg --range 6 to send reminders for birthdays within next 7 days, \nnote that remainder will be for each birthday and default range is 0 to 6 default is 0")
}
