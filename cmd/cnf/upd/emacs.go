package upd

import (
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var UpdCnfEmacsCmd = &cobra.Command{
	Use: "emacs",
	Short: "CLI to update emacs config",
	Run: func(cmd *cobra.Command, args []string) {
		command.CreateEmacsInit()
	},
}
