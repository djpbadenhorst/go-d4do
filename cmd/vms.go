package cmd

import (
	"log"
	"d4do/pkg/install"
	"d4do/pkg/command"
	"github.com/spf13/cobra"
)

var VmsCmd = &cobra.Command{
	Use:   "vms",
	Short: "List Virtual Machines",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Vagrant")
		vms := command.VagrantListVms()
		log.Println(vms)
	},
}
