package main

import (
	"identity-web-api/server"
	"identity-web-api/storage"
	"log"
)

func main() {
	var setting = GetAppSetting()
	var serverConfig = server.Config{
		IP:   setting.ServerIp,
		Port: setting.ServerPort,
	}
	var connectionString = setting.DbConnectionString
	var storageConfig = storage.Config{ConnectionString: connectionString}
	var dataStorage = storage.CreatePQStore(&storageConfig)
	var s = server.Create(serverConfig, dataStorage)
	var err = s.Start()
	log.Fatalln(err)
}
