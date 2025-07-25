package common

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UserCors(engine *gin.Engine) {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AddAllowHeaders("Authorization")
	engine.Use(cors.New(cfg))
}
