package cmd

import (
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var UpgCmd = &cobra.Command{
	Use: "upg",
	Short: "CLI to run apt update and upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		command.Exec("sudo", "apt", "update")
		command.Exec("sudo", "apt", "upgrade", "-y")
	},
}
