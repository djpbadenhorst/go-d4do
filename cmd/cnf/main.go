package cnf

import (
	updCli "d4do/cmd/cnf/upd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "cnf",
	Short: "CLI to manage application configs",
}

func init() {
	RootCmd.AddCommand(updCli.RootCmd)
}
