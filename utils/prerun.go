package utils

import "github.com/spf13/cobra"

func ValidFlags(cmd *cobra.Command, args []string) {
	name, errn := cmd.Flags().GetString("name")
	month, errm := cmd.Flags().GetString("month")
	date, errd := cmd.Flags().GetString("date")
	mobile, errmo := cmd.Flags().GetString("phone")

	if errd != nil || errmo != nil || errm != nil || errn != nil {
		panic("check the flags, flags missing")
	}
	if len(name) == 0 || len(mobile) == 0 || len(month) == 0 || len(date) == 0 {
		panic("all flags are required")
	}
}

func ValidFlagsEdit(cmd *cobra.Command, args []string) {
	name, errn := cmd.Flags().GetString("name")
	mobile, errmo := cmd.Flags().GetString("mobile")

	if errmo != nil || errn != nil {
		panic("check the flags, flags missing")
	}
	if len(name) == 0 || len(mobile) == 0 {
		panic("all flags are required")
	}
}
