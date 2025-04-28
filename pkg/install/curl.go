package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func CurlInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.Exec("sudo", "apt", "install", "-y", "curl")
	default:
		log.Fatalln("Cannot install Curl on current operating system: ", opSys)
	}
}

func CurlCheck() bool {
	if err := exec.Command("curl", "--version").Run(); err==nil {
		return true
	} 
	return false
}
