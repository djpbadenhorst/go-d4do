package command

import (
	"os"
	"d4do/pkg/config"
)

func CreateKittyConfig() {
	CreateDirIfNotExists(os.Getenv("HOME") + "/.config")
	CreateDirIfNotExists(os.Getenv("HOME") + "/.config/kitty")
	os.WriteFile(os.Getenv("HOME") + "/.config/kitty/kitty.conf", []byte(config.KittyConfig), os.ModePerm)
}
