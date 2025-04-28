package config

import (
	"os"
	"log"
//	yaml "gopkg.in/yaml.v3"
)

var AppConfigDir = os.Getenv("HOME") + "/.d4do"
var AppConfigFilePath = AppConfigDir + "/config.yaml"

func InitConfig() {
	if _, err1 := os.Stat(AppConfigDir); os.IsNotExist(err1) {
		if err2 := os.Mkdir(AppConfigDir, os.ModePerm); err2 != nil {
			log.Fatal(err2)
		}
		log.Println("Creating directory:", AppConfigDir)
	}

	if _, err1 := os.Stat(AppConfigFilePath); os.IsNotExist(err1) {
		if err2 := os.WriteFile(AppConfigFilePath, []byte(""), os.ModePerm); err2 != nil {
			log.Fatal(err2)
		}
		log.Println("Creating file:", AppConfigFilePath)
	}
}

//type ConfigPackageStruct struct {
//	Name string `mapstructure:"name"`
//}

//type ConfigStruct struct {
//	//Packages []ConfigPackageStruct `mapstructure:"packages"`
//}
//var AppConfig ConfigStruct

//func InitConfigDep() {
//	command.CreateDirIfNotExists(AppConfigDir)
//	command.CreateFileIfNotExists(AppConfigFilePath)
//	command.CreateDirIfNotExists(AppConfigDir + "/vagrant")
//
//	yamlBytes, err := os.ReadFile(AppConfigFile)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if err := yaml.Unmarshal(yamlBytes, &AppConfig); err != nil {
//		log.Fatal(err)
//	}
//	log.Println("Read config from file:", AppConfigFile)
//
//	SaveConfig()
//}

//func SaveConfig() {
//	yamlBytes, err := yaml.Marshal(AppConfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//	
//	err = os.WriteFile(AppConfigFile, yamlBytes, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
