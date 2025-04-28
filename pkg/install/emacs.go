package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func EmacsInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		InstallIfNotInstalled("Kitty")
		command.CreateEmacsInit()
		command.Exec("sudo", "apt", "install", "-y", "emacs-nox")
	default:
		log.Fatalln("Cannot install Emacs on current operating system: ", opSys)
	}
}

func EmacsCheck() bool {
	if err := exec.Command("emacs", "--version").Run(); err==nil {
		return true
	} 
	return false
}
