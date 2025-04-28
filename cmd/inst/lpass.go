package inst

import (
	"d4do/pkg/install"
	"github.com/spf13/cobra"
)

var InstLastPassCmd = &cobra.Command{
	Use:   "lpass",
	Short: "Install LastPass with gvm",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("LastPass")
	},
}
