package install

import (
	"os"
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func GolangInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.Exec("sudo", "apt", "install", "-y", "libfuse-dev")
		command.Exec("sudo", "apt", "install", "-y", "gcc")
		command.Exec("sudo", "apt", "install", "-y", "make")
		command.Exec("sudo", "apt", "install", "-y", "bison")
		command.Exec("bash", "-c", "curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash")

		cmdPrefix := "source " + os.Getenv("HOME") + "/.gvm/scripts/gvm"
		command.Exec("bash", "-c", cmdPrefix + "; gvm install go1.4 -B")
		command.Exec("bash", "-c", cmdPrefix + "; gvm use go1.4; gvm install go1.17.13")
		command.Exec("bash", "-c", cmdPrefix + "; gvm use go1.17.13; gvm install go1.21.13")
	default:
		log.Fatalln("Cannot install Golang on current operating system: ", opSys)
	}
}

func GolangCheck() bool {
	if err := exec.Command("gvm", "version").Run(); err==nil {
		return true
	} 
	return false
}
