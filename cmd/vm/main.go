package vm

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "vm",
	Short: "CLI to manage Virtual Machines",
}

func init() {	
	RootCmd.AddCommand(VmUpCmd)
	RootCmd.AddCommand(VmDownCmd)
	RootCmd.AddCommand(VmStartCmd)
	RootCmd.AddCommand(VmStopCmd)
	RootCmd.AddCommand(VmSshCmd)
	RootCmd.PersistentFlags().StringP("vm-name", "n", "default", "Name of the machine")
	RootCmd.MarkFlagRequired("vm-name")
}
