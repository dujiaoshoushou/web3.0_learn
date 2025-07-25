package handlers

import (
	"github.com/gin-gonic/gin"
	"task-week-four/handlers/comment"
	"task-week-four/handlers/common"
	"task-week-four/handlers/post"
	"task-week-four/handlers/role"
	"task-week-four/handlers/system"
	"task-week-four/handlers/user"
)

func InitEngine() *gin.Engine {
	r := gin.Default()
	common.UserCors(r)
	// 注册JWT用户认证鉴权中间件
	common.UseJWTAuth(r)
	user.Router(r)
	system.Router(r)
	role.Router(r)
	post.Router(r)
	comment.Router(r)

	return r
}
