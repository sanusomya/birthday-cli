package cmd

import (
	get "github.com/sanusomya/birthday-cli/cmd/get"
	add "github.com/sanusomya/birthday-cli/cmd/add"
	edit "github.com/sanusomya/birthday-cli/cmd/edit"
	delete "github.com/sanusomya/birthday-cli/cmd/delete"
	"github.com/sanusomya/birthday-cli/config"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:               "birthday",
	Aliases:           []string{"birth", "bday"},
	SuggestFor:        []string{"birtg", "b", "d", "bay", "irth"},
	Short:             "use this command with flags and args to get started",
	Example:           "birthday <args> -<flag> <data>",
	Version:           config.Version,
	Run:               runRoot,
	DisableAutoGenTag: true,
}

func runRoot(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("you need an argument first. try -h")
		os.Exit(1)
	}
}

func init() {
	Root.AddCommand(get.CmdGet)
	Root.AddCommand(add.CmdAdd)
	Root.AddCommand(delete.CmdDelete)
	Root.AddCommand(edit.CmdEdit)
}
