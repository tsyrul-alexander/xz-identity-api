package main

import (
	"./server"
	"fmt"
)

func main() {
	var setting = GetAppSetting()
	fmt.Println(setting.Storage)
	var serverConfig = server.Config{
		Ip:   "localhost",
		Port: 8080,
	}
	var s = server.Server{Config:serverConfig}
	var err = s.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
