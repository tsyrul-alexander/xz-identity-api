package main

import (
	"./server"
	"./storage"
	"fmt"
)

func main() {
	var setting = GetAppSetting()
	fmt.Println(setting.Storage)
	var serverConfig = server.Config{
		IP:   "localhost",
		Port: 8080,
	}
	var connectionString = setting.Storage["connectionString"].(string)
	var storageConfig = storage.Config{ConnectionString:connectionString}
	var dataStorage = storage.CreatePQStore(&storageConfig)
	var s = server.Create(serverConfig, dataStorage)
	var err = s.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
