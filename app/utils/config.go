package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"weapp/app/config"
)

func ReadSettingsFromFile(settingFilePath string) (config config.Config) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
