package post

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//分支路由
	g := r.Group("post")
	g.POST("create", Create)
	g.POST("list", GetList)
	g.POST("update/:id", EditPost)
	g.POST("delete/:id", DeletePost)
	g.GET("", GetRaw)

}
