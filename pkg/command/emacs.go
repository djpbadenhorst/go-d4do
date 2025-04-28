package command

import (
	"os"
	"d4do/pkg/config"
)

func CreateEmacsInit() {
	CreateDirIfNotExists(os.Getenv("HOME") + "/.emacs.d")
	os.WriteFile(os.Getenv("HOME") + "/.emacs.d/init.el", []byte(config.EmacsConfig), os.ModePerm)
}
