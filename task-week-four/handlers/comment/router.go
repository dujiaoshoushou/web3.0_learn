package comment

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	//分支路由
	g := r.Group("comment")
	g.POST("create", Create)
	g.POST("list", GetList)

}
