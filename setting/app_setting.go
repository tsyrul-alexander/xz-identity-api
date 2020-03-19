package setting

import (
	"github.com/spf13/viper"
	"log"
)

type AppSetting struct {
	DbConnectionString string `mapstructure:"DB_CONNECTION_STRING"`
	ServerIp string `mapstructure:"SERVICE_IP"`
	ServerPort int `mapstructure:"SERVER_PORT"`
	JwtKey string `mapstructure:"JWT_KEY"`
}
const FilePath string = "config.json"
var instance *AppSetting

func GetAppSetting() *AppSetting {
	if instance == nil {
		instance = getAppConfig()
	}
	return instance
}

func getAppConfig() *AppSetting {
	var config = AppSetting{}
	var v = configureViper()
	setSettingValue(v, &config)
	return &config
}
func setSettingValue(v *viper.Viper, setting *AppSetting) {
	if err := v.Unmarshal(&setting); err != nil {
		log.Fatalln(err.Error())
	}
}
func configureViper() *viper.Viper {
	var v = viper.New()
	setDefaultValues(v)
	setConfigFile(v)
	v.AutomaticEnv()
	return v
}
func setConfigFile(v *viper.Viper)  {
	v.SetConfigFile(FilePath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err.Error())
	}
}
func setDefaultValues(v *viper.Viper) {
	var defaultSettings = getDefaultSettingValues()
	for key, value := range defaultSettings {
		v.SetDefault(key, value)
	}
}

func getDefaultSettingValues() map[string]interface{} {
	return map[string]interface{} {
		"SERVICE_IP": "0.0.0.0",
		"SERVER_PORT": "8080",
	}
}