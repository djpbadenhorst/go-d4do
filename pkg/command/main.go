package command

import (
	"os"
	"log"
	"strings"
	"os/exec"
)

func Exec(entrypoint string, command ...string) {
	log.Println("Running command:", entrypoint, strings.Join(command, " "))
	shCmd := exec.Command(entrypoint, command...)
	
	shCmd.Stdin = os.Stdin
	shCmd.Stdout = os.Stdout
	
	if err := shCmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func CreateDirIfNotExists(dir string) {
	if _, err1 := os.Stat(dir); os.IsNotExist(err1) {
		if err2 := os.Mkdir(dir, os.ModePerm); err2 != nil {
			log.Fatal(err2)
		}
		log.Println("Creating directory:", dir)
	}
}

func CreateFileIfNotExists(filePath string) {
	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		if err2 := os.WriteFile(filePath, []byte(""), os.ModePerm); err2 != nil {
			log.Fatal(err2)
		}
		log.Println("Creating file:", filePath)
	}
}
