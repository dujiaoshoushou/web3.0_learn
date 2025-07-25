package system

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/ping", Ping)
}
