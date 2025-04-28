package upd

import (
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var UpdCnfZshCmd = &cobra.Command{
	Use: "zsh",
	Short: "CLI to update zsh config",
	Run: func(cmd *cobra.Command, args []string) {
		command.CreateZshConfig()
	},
}
