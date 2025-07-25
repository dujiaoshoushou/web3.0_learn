package utils

import (
	"github.com/spf13/viper"
	"log"
)

func defaultConfig() {
	viper.SetDefault("app.addr", ":8080")
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("app.log.path", "./logs")
	viper.SetDefault("db.dsn", "root:123456@tcp(127.0.0.1:3307)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local")
}

func ParseConfig() {
	// 1. 默认配置
	defaultConfig()
	// 2. 配置解析参数
	viper.AddConfigPath(".")
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	// 3. 执行解析
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed, err: " + err.Error())
	}
}
