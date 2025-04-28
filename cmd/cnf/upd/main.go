package upd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "upd",
	Short: "CLI to update application configs",
}

func init() {
	RootCmd.AddCommand(UpdCnfEmacsCmd)
	RootCmd.AddCommand(UpdCnfZshCmd)
}
