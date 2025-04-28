package install

import (
	"log"
)

func InstallIfNotInstalled(pkgName string) {
	switch pkgName {
	case "Emacs":
		if isInstalled := EmacsCheck(); !isInstalled {
			log.Println("Installing Emacs")
			EmacsInstall()
		}
	case "Golang":
		if isInstalled := GolangCheck(); !isInstalled {
			log.Println("Installing Golang")
			GolangInstall()
		}
	case "LastPass":
		if isInstalled := LastPassCheck(); !isInstalled {
			log.Println("Installing LastPass")
			LastPassInstall()
		}
	case "Zsh":
		if isInstalled := ZshCheck(); !isInstalled {
			log.Println("Installing Zsh")
			ZshInstall()
		}
	case "Curl":
		if isInstalled := CurlCheck(); !isInstalled {
			log.Println("Installing Curl")
			CurlInstall()
		}
	case "Eza":
		if isInstalled := EzaCheck(); !isInstalled {
			log.Println("Installing Eza")
			EzaInstall()
		}
	case "Tofu":
		if isInstalled := TofuCheck(); !isInstalled {
			log.Println("Installing Tofu")
			TofuInstall()
		}
	case "Kitty":
		if isInstalled := KittyCheck(); !isInstalled {
			log.Println("Installing Kitty")
			KittyInstall()
		}
	case "Unzip":
		if isInstalled := UnzipCheck(); !isInstalled {
			log.Println("Installing Unzip")
			UnzipInstall()
		}
	case "Vagrant":
		if isInstalled := VagrantCheck(); !isInstalled {
			log.Println("Installing Vagrant")
			VagrantInstall()
		}
	default:
		log.Fatalln("No installation instructions for: ", pkgName)
	}
}
