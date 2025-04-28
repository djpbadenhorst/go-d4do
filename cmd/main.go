package cmd

import (
	"d4do/pkg/config"
	vmCli "d4do/cmd/vm"
	cnfCli "d4do/cmd/cnf"
	instCli "d4do/cmd/inst"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "d4do",
	Short: "Command line interface utility",
}

func Run() {
	RootCmd.Execute()
}

func init() {
	config.InitConfig()
	
	RootCmd.CompletionOptions.DisableDefaultCmd = false
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	
	RootCmd.AddCommand(instCli.RootCmd)
	RootCmd.AddCommand(cnfCli.RootCmd)
	RootCmd.AddCommand(VmsCmd)
	RootCmd.AddCommand(UpgCmd)
	RootCmd.AddCommand(vmCli.RootCmd)
}
