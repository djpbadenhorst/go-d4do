package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func UnzipInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.Exec("sudo", "apt", "install", "-y", "unzip")
	default:
		log.Fatalln("Cannot install Unzip on current operating system: ", opSys)
	}
}

func UnzipCheck() bool {
	if err := exec.Command("unzip", "--version").Run(); err==nil {
		return true
	} 
	return false
}
