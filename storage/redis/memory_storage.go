package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/identity-web-api/model/memory"
	"log"
)

type MemoryStorage struct {
	client *redis.Client
}

const UserKeyFormat = "User_%v"

func Create(address string, password string, db int) *MemoryStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &MemoryStorage{client:client}
}

func (ms *MemoryStorage)SetUser(user *memory.User) bool {
	var key = getUserKey(user.Id)
	return ms.setData(key, user)
}

func (ms *MemoryStorage)GetUser(id uuid.UUID) (user *memory.User, ifExists bool)  {
	var key = getUserKey(id)
	var u memory.User
	return &u, ms.getData(key, &u)
}

func (ms *MemoryStorage)getData(key string, data interface{}) bool {
	var dataStr = ms.client.Get(key)
	var value, err = dataStr.Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		log.Println(err.Error())
	}
	parseStr(value, data)
	return true
}

func (ms *MemoryStorage)setData(key string, data interface{}) bool {
	var dataStr = getStr(data)
	var status = ms.client.Set(key, dataStr, 0)
	var err = status.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return err != nil
}

func getUserKey(id uuid.UUID) string {
	return fmt.Sprintf(UserKeyFormat, id.String())
}

func getStr(data interface{}) string {
	var bytes, err = json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}

func parseStr(str string, data interface{}) {
	if err := json.Unmarshal([]byte(str), data); err != nil {
		log.Println(err.Error())
	}
}
