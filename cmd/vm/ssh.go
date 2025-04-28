package vm

import (
	"d4do/pkg/install"
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var VmSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into Virtual Machine",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")
		command.VagrantCommand("vagrant ssh", cmd.Flag("vm-name").Value.String())
	},
}
