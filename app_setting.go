package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type AppSetting struct {
	Storage map[string]interface{} `json:"storage"`
}

const settingFileName = "app.json"
var instance *AppSetting

func GetAppSetting() *AppSetting {
	if instance == nil {
		instance, _ = getAppConfig()
	}
	return instance
}

func getAppConfig() (*AppSetting, error) {
	jsonFile, err := os.Open(settingFileName)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var setting AppSetting
	err = json.Unmarshal(bytes, &setting)
	if err != nil {
		return nil, err
	}
	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}
	return &setting, nil
}