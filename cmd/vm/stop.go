package vm

import (
	"d4do/pkg/install"
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var VmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Virtual Machine",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")
		command.VagrantCommand("vagrant halt", cmd.Flag("vm-name").Value.String())
	},
}
