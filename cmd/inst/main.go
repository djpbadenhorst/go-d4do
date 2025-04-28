package inst

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "inst",
	Short: "CLI to manage installations",
}

func init() {
	RootCmd.AddCommand(InstGolangCmd)
	RootCmd.AddCommand(InstZshCmd)
	RootCmd.AddCommand(InstLastPassCmd)
	RootCmd.AddCommand(InstEmacsCmd)
	RootCmd.AddCommand(InstTofuCmd)
}
