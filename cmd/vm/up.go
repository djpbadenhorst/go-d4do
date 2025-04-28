package vm

import (
	"d4do/pkg/config"
	"d4do/pkg/install"
	"d4do/pkg/command"	
	"github.com/spf13/cobra"

	"d4do/pkg/tui/passwordPrompt"
)

var VmUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Create Virtual Machine",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")
		command.CreateDirIfNotExists(config.AppConfigDir + "/vagrant/" + cmd.Flag("vm-name").Value.String())
		
		password := passwordPrompt.Run()
		command.VagrantCommandWithPassword("vagrant up", cmd.Flag("vm-name").Value.String(), password)
	},
}
