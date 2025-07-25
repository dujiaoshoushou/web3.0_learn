package user

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//分支路由
	g := r.Group("user")
	g.POST("login", Login)
	g.POST("regist", Regist)

}
