package utils

import(
	"github.com/spf13/cobra"
)
func Exec(c *cobra.Command) error{
	return c.Execute()
}