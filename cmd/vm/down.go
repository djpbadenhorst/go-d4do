package vm

import (
	"d4do/pkg/install"
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var VmDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Destroy Virtual Machine",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")
		command.VagrantCommand("vagrant destroy", cmd.Flag("vm-name").Value.String())
	},
}
