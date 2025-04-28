package install

import (
	"log"
	"runtime"
	"os/exec"
	"d4do/pkg/command"
)

func ZshInstall() {
	switch opSys := runtime.GOOS; opSys {
	case "linux":
                InstallIfNotInstalled("Kitty")
		InstallIfNotInstalled("Eza")
		
		command.Exec("sudo", "apt", "install", "-y", "zsh")
		command.Exec("bash", "-c", "curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh | sh")
		command.CreateZshConfig()
	default:
		log.Fatalln("Cannot install Zsh on current operating system: ", opSys)
	}
}

func ZshCheck() bool {
	if err := exec.Command("zsh", "--version").Run(); err==nil {
		return true
	} 
	return false
}
