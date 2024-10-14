package cmd

import (
	"fmt"
	"os"

	add "github.com/sanusomya/birthday-cli/cmd/add"
	delete "github.com/sanusomya/birthday-cli/cmd/delete"
	edit "github.com/sanusomya/birthday-cli/cmd/edit"
	get "github.com/sanusomya/birthday-cli/cmd/get"
	"github.com/sanusomya/birthday-cli/cmd/wish"
	"github.com/sanusomya/birthday-cli/config"
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
	Root.AddCommand(wish.CmdWish)
}
