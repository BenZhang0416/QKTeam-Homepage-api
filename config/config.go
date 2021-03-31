package config

import (
	"api/boot/config"
	"fmt"
	"time"

	"github.com/spf13/viper" //读取服务器配置的包
)

func init() {
	fmt.Println("init config")
	config.SetCheckDuration(time.Minute)
	InitHttpConfig()
}

func HttpConfig() *httpConf {
	return httpConfig.GetConfig().(*httpConf)
}

func GetString(name string) string {
	return viper.GetString(name)
}
