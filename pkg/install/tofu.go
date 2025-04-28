package install

import (
	"os"
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/config"
	"d4do/pkg/command"
)

func TofuInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		tempDir, _ := os.MkdirTemp(config.AppConfigDir, "temp")
		defer os.RemoveAll(tempDir)

		url := "https://github.com/opentofu/opentofu/releases/download/v1.9.1/tofu_1.9.1_linux_amd64.tar.gz"
		command.Exec("wget", url, "-O", tempDir + "/tofu.tar.gz")
		command.Exec("tar", "-xvf", tempDir + "/tofu.tar.gz", "-C", tempDir)
		command.Exec("sudo", "mv", tempDir + "/tofu", "/usr/bin/")

	default:
		log.Fatalln("Cannot install Tofu on current operating system: ", opSys)
	}
}

func TofuCheck() bool {
	if err := exec.Command("tofu", "--version").Run(); err==nil {
		return true
	} 
	return false
}
