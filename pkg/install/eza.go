package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func EzaInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.Exec("sudo", "apt", "install", "-y", "eza")
	default:
		log.Fatalln("Cannot install Eza on current operating system: ", opSys)
	}
}

func EzaCheck() bool {
	if err := exec.Command("eza", "--version").Run(); err==nil {
		return true
	} 
	return false
}
