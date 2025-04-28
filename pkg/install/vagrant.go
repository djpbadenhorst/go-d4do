package install

import (
	"os"
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/config"
	"d4do/pkg/command"
)

func VagrantInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		tempDir, _ := os.MkdirTemp(config.AppConfigDir, "temp")
		defer os.RemoveAll(tempDir)

		if isInstalled := CurlCheck(); !isInstalled {
			log.Println("Installing Curl")
			CurlInstall()
		}

		if isInstalled := UnzipCheck(); !isInstalled {
			log.Println("Installing Unzip")
			UnzipInstall()
		}

		command.Exec("sudo", "apt", "install", "-y", "libfuse-dev")
		command.Exec("curl", "-o", tempDir + "/vagrant.zip", "https://releases.hashicorp.com/vagrant/2.4.3/vagrant_2.4.3_linux_amd64.zip")
		command.Exec("unzip", tempDir + "/vagrant.zip", "-d", tempDir)
		command.Exec("sudo", "mv", tempDir + "/vagrant", "/usr/bin/")
		command.Exec("sudo", "apt", "install", "-y", "virtualbox")

		command.CreateDirIfNotExists(config.AppConfigDir + "/vagrant")
	default:
		log.Fatalln("Cannot install Vagrant on current operating system: ", opSys)
	}
}

func VagrantCheck() bool {
	if err := exec.Command("vagrant", "--version").Run(); err==nil {
		return true
	} 
	return false
}
