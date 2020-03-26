package main

import (
	"github.com/tsyrul-alexander/identity-web-api/server"
	"github.com/tsyrul-alexander/identity-web-api/setting"
	"github.com/tsyrul-alexander/identity-web-api/storage"
	"github.com/tsyrul-alexander/identity-web-api/storage/pq"
	"github.com/tsyrul-alexander/identity-web-api/storage/redis"
	"log"
)

func main() {
	var dataStorage = ConfigureDataStorage()
	var memoryStorage = ConfigureMemoryStorage()
	var s = ConfigureServer(dataStorage, memoryStorage)
	var err = s.Start()
	log.Fatalln(err)
}

func ConfigureServer(dataStorage storage.DataStorage, memoryStorage storage.MemoryStorage) *server.Server {
	var st = setting.GetAppSetting()
	var serverConfig = server.Config {
		IP:   st.Server.Ip,
		Port: st.Server.Port,
	}
	return server.Create(serverConfig, dataStorage, memoryStorage)
}

func ConfigureDataStorage() *pq.DataStorage {
	var st = setting.GetAppSetting()
	var connectionString = st.Storage.Data.PQ.ConnectionString
	var storageConfig = storage.Config{ConnectionString: connectionString}
	return pq.CreateStore(&storageConfig)
}

func ConfigureMemoryStorage() *redis.MemoryStorage {
	var st = setting.GetAppSetting()
	var redisSetting = st.Storage.Memory.Redis
	return redis.Create(redisSetting.Address, redisSetting.Password, redisSetting.Db)
}