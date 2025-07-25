package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-week-four/models"
	"task-week-four/utils"
)

func Create(c *gin.Context) {
	req := &AddCommentReq{}
	if err := c.ShouldBind(req); err != nil {
		utils.Logger().Error(err.Error(), "status", req.Status, "content", req.Content)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "创建文章评论错误.",
		})
		return
	}
	d, _ := c.Get("user_id")
	println(d)
	user_id := uint64(d.(float64))
	req.UserID = &user_id
	row := req.ToPost()
	if err := models.CommentInsert(row); err != nil {
		utils.Logger().Error(err.Error(), "status", req.Status, "content", req.Content)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "创建文章评论错误.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": row,
		"msg":  "创建文章评论成功",
	})

}

func GetList(c *gin.Context) {
	list(c, models.SCOPE_UNDELETED, false)
}

func list(c *gin.Context, scope uint8, assoc bool) {
	req := GetListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章评论错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章评论错误",
		})
		return
	}
	//log.Println(req)
	log.Println(req.Content, req.PageNum, req.PageSize, req.SortField, req.SortMethod)
	req.Clean()

	rows, total, err := models.CommentFecthList(assoc, req.CommentFilter, req.Sorter, req.Page, scope)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章评论错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章评论错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  rows,
		"total": total,
		"msg":   "success",
	})

}
