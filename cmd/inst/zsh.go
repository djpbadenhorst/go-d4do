package inst

import (
	"d4do/pkg/install"
	"github.com/spf13/cobra"
)

var InstZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Install Zsh with gvm",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Zsh")
	},
}
