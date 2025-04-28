package inst

import (
	"d4do/pkg/install"
	"github.com/spf13/cobra"
)

var InstGolangCmd = &cobra.Command{
	Use:   "go",
	Short: "Install Golang with gvm",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Golang")
	},
}
