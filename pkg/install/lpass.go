package install

import (
	"os"
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func LastPassInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
		command.CreateDirIfNotExists(os.Getenv("HOME") + "/.config/lpass")
		command.CreateDirIfNotExists(os.Getenv("HOME") + "/.local")
		command.CreateDirIfNotExists(os.Getenv("HOME") + "/.local/share")
		command.CreateDirIfNotExists(os.Getenv("HOME") + "/.local/share/lpass")
		command.Exec("sudo", "apt", "install", "-y", "lastpass-cli")
	default:
		log.Fatalln("Cannot install LastPass on current operating system: ", opSys)
	}
}

func LastPassCheck() bool {
	if err := exec.Command("lpass", "--version").Run(); err==nil {
		return true
	} 
	return false
}
