package inst

import (
	"d4do/pkg/install"
	"github.com/spf13/cobra"
)

var InstTofuCmd = &cobra.Command{
	Use:   "tofu",
	Short: "Install Tofu with gvm",
	Run: func(cmd *cobra.Command, args []string) {
		install.InstallIfNotInstalled("Tofu")
	},
}
