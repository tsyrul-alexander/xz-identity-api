package main

import (
	"github.com/tsyrul-alexander/identity-web-api/server"
	"github.com/tsyrul-alexander/identity-web-api/setting"
	"github.com/tsyrul-alexander/identity-web-api/storage"
	"github.com/tsyrul-alexander/identity-web-api/storage/pq"
	"log"
)

func main() {
	var setting = setting.GetAppSetting()
	var serverConfig = server.Config{
		IP:   setting.ServerIp,
		Port: setting.ServerPort,
	}
	var connectionString = setting.DbConnectionString
	var storageConfig = storage.Config{ConnectionString: connectionString}
	var dataStorage = pq.CreateStore(&storageConfig)
	var s = server.Create(serverConfig, dataStorage)
	var err = s.Start()
	log.Fatalln(err)
}
