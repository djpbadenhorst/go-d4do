package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func KittyInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.CreateKittyConfig()
		command.Exec("sudo", "apt", "install", "-y", "kitty")
	default:
		log.Fatalln("Cannot install Kitty on current operating system: ", opSys)
	}
}

func KittyCheck() bool {
	if err := exec.Command("kitty", "--version").Run(); err==nil {
		return true
	} 
	return false
}
