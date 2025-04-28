package inst

import (
	"d4do/pkg/install"
	"github.com/spf13/cobra"
)

var InstEmacsCmd = &cobra.Command{
	Use:   "emacs",
	Short: "Install Emacs with gvm",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Emacs")
	},
}
