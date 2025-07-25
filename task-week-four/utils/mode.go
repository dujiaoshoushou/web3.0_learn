package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

// 定义应用三种模式
const (
	APP_DEBUG   = "DEBUG"
	APP_RELEASE = "RELEASE"
	APP_TEST    = "TEST"
)

// 设置应用的启动模式
func SetMode() {
	switch strings.ToUpper(viper.GetString("app.mode")) {
	case APP_DEBUG:
		gin.SetMode(gin.DebugMode)
	case APP_RELEASE:
		gin.SetMode(gin.ReleaseMode)
	case APP_TEST:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
