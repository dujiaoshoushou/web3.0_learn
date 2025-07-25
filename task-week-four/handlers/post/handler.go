package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-week-four/models"
	"task-week-four/utils"
)

func Create(c *gin.Context) {
	req := &AddPostReq{}
	if err := c.ShouldBind(req); err != nil {
		utils.Logger().Error(err.Error(), "title", req.Title, "content", req.Content)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "创建文章错误.",
		})
		return
	}
	d, _ := c.Get("user_id")
	println(d)
	user_id := uint64(d.(float64))
	req.UserID = &user_id
	row := req.ToPost()
	if err := models.PostInsert(row); err != nil {
		utils.Logger().Error(err.Error(), "title", req.Title, "content", req.Content)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "创建文章错误.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": row,
		"msg":  "创建文章成功.",
	})

}

func GetList(c *gin.Context) {
	list(c, models.SCOPE_UNDELETED, false)
}

func EditPost(c *gin.Context) {
	req := &EditReq{}
	if err := c.ShouldBindUri(req); err != nil {
		utils.Logger().Error(err.Error(), "message", "编辑文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "编辑文章错误",
		})
		return
	}
	user_id, _ := c.Get("user_id")
	userID := uint64(user_id.(float64))
	row, err := models.PostFetchRow(false, "id = ? and user_id = ?", req.ID, userID)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "当前用户没有编辑此文章权限", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "当前用户没有编辑此文章权限",
		})
		return
	}
	if row == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msge": "当前用户没有编辑此文章权限",
		})
		return
	}
	reqBody := &EditBodyReq{}
	if err := c.ShouldBindBodyWithJSON(reqBody); err != nil {
		utils.Logger().Error(err.Error(), "message", "编辑文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "编辑文章错误",
		})
		return
	}

	fieldMap := reqBody.ToFiledMap()
	rows, err := models.PostEdit(fieldMap, req.ID)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "编辑文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "编辑文章错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rows,
		"msg":  "success",
	})
}

func DeletePost(c *gin.Context) {
	req := &DeleteReq{}
	if err := c.ShouldBindUri(req); err != nil {
		utils.Logger().Error(err.Error(), "message", "删除文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "删除文章错误",
		})
		return
	}
	user_id, _ := c.Get("user_id")
	userID := uint64(user_id.(float64))
	row, err := models.PostFetchRow(false, "id =? and user_id =?", req.ID, userID)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "当前用户没有删除此文章权限", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "当前用户没有删除此文章权限",
		})
		return
	}
	if row == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msge": "当前用户没有删除此文章权限",
		})
		return
	}
	rows, err := models.PostDelete(req.ID, false)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "删除文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "删除文章错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rows,
		"msg":  "删除文件成功",
	})

}

func GetRaw(c *gin.Context) {
	req := GetRowReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章错误",
		})
		return
	}
	row, err := models.PostFetchRow(false, "title like ?", "%"+req.Title+"%")
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": row,
		"msg":  "success",
	})
	utils.Logger().Info("角色查询单条result: ", row)
}

func list(c *gin.Context, scope uint8, assoc bool) {
	req := GetListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章错误",
		})
		return
	}
	//log.Println(req)
	log.Println(req.Title, req.PageNum, req.PageSize, req.SortField, req.SortMethod)
	req.Clean()

	rows, total, err := models.PostFecthList(assoc, req.PostFilter, req.Sorter, req.Page, scope)
	if err != nil {
		utils.Logger().Error(err.Error(), "message", "查询文章错误", "req", req)
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询文章错误",
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
