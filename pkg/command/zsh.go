package command

import (
	"os"
	"d4do/pkg/config"
)

func CreateZshConfig() {
	dir1 := os.Getenv("HOME") + "/powerlevel10k"
	if _, err1 := os.Stat(dir1); os.IsNotExist(err1) {
		Exec("git", "clone", "--depth=1", "https://github.com/romkatv/powerlevel10k.git", dir1)
	}

	dir2 := os.Getenv("HOME") + "/.oh-my-zsh/custom/plugins/zsh-autosuggestions"
	if _, err1 := os.Stat(dir2); os.IsNotExist(err1) {
		Exec("git", "clone", "https://github.com/zsh-users/zsh-autosuggestions", dir2)
	}

	dir3 := os.Getenv("HOME") + "/zsh-syntax-highlighting/"
	if _, err1 := os.Stat(dir3); os.IsNotExist(err1) {
		Exec("git", "clone", "https://github.com/zsh-users/zsh-syntax-highlighting.git", dir3)
	}

	os.WriteFile(os.Getenv("HOME") + "/.zshrc", []byte(config.ZshConfig), os.ModePerm)
	os.WriteFile(os.Getenv("HOME") + "/.p10k.zsh", []byte(config.P10kConfig), os.ModePerm)
}
