package vm

import (
	"log"
	"slices"
	"d4do/pkg/install"
	"d4do/pkg/command"	
	"github.com/spf13/cobra"
)

var VmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Virtual Machine",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")

		vmName := cmd.Flag("vm-name").Value.String()
		vms := command.VagrantListVms()
		
		if slices.Contains(vms, vmName) {
			command.VagrantCommand("vagrant up", vmName)
		} else {
			log.Fatalln("No vm with name:", vmName)
		}
	},
}
