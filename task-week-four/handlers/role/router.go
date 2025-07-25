package role

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Router(r *gin.Engine) {
	//分支路由
	g := r.Group("role")
	g.GET("", GetRaw)
	g.GET("list", GetList)
	g.POST("add", Add)
	g.DELETE("delete", Delete)
	g.GET("recycle", Recycle)
	g.PUT("restore", Restore)
	g.PUT("update/:id", Edit)
	g.PUT("update_batch_enable", EditBatchEnable)
}

func init() {
	registerValidator()
}

func registerValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("roleTitleUnique", roleTitleUnique)
	}
}
