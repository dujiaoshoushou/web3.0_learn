package role

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-week-four/handlers/common"
	"task-week-four/models"
	"task-week-four/utils"
)

// GetRaw 获取角色单条
// @Summary 获取角色单条
// @Description 获取角色单条
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body GetRowReq true "请求体"
// @Success 200 {object} GetRowRes
// @Router /role/raw [get]
func GetRaw(c *gin.Context) {
	req := GetRowReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	row, err := models.RoleFetchRow(false, "id = ?", req.ID)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
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

// GetList 获取角色列表
// @Summary 获取角色列表
// @Description 获取角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body GetListReq true "请求体"
// @Success 200 {object} GetListRes
// @Router /role/list [get]
func GetList(c *gin.Context) {
	//req := GetListReq{}
	//if err := c.ShouldBindQuery(&req); err != nil {
	//	utils.Logger().Error(err.Error())
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 100,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	////log.Println(req)
	//log.Println(req.Keyword, req.PageNum, req.PageSize, req.SortField, req.SortMethod)
	//req.Clean()
	//
	//rows, total, err := models.RoleFecthList(false, req.RoleFilter, req.Sorter, req.Page, models.SCOPE_UNDELETED)
	//if err != nil {
	//	utils.Logger().Error(err.Error())
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 100,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code":  0,
	//	"data":  rows,
	//	"total": total,
	//	"msg":   "success",
	//})
	list(c, models.SCOPE_UNDELETED, false)
}

// Add 添加角色
// @Summary 添加角色
// @Description 添加角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body AddReq true "请求体"
// @Success 200 {object} AddRes
// @Router /role/add [post]
func Add(c *gin.Context) {
	req := AddReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":         100,
			"msg":          err.Error(),
			"transMessage": common.Translate(err),
		})
		return
	}
	log.Println(req)
	role := req.ToRole()
	if err := models.RoleInsert(role); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "数据插入失败",
		})
		return
	}

	row, err := models.RoleFetchRow(false, "id = ?", role.ID)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "查询错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": row,
		"msg":  "success",
	})
}

// Delete 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body DeleteReq true "请求体"
// @Success 200 {object} DeleteRes
// @Router /role/delete [delete]
func Delete(c *gin.Context) {
	req := DeleteReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(req)

	rosNum, err := models.RoleDelete(req.IDList, req.Force)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rosNum,
		"msg":  "success",
	})
}

func Recycle(c *gin.Context) {
	list(c, models.SCOPE_DELETED, false)
}

// list 列表
// @Summary 列表
// @Description 列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body GetListReq true "请求体"
// @Success 200 {object} GetListRes
// @Router /role/list [get]
func list(c *gin.Context, scope uint8, assoc bool) {
	req := GetListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(req)
	log.Println(req.Keyword, req.PageNum, req.PageSize, req.SortField, req.SortMethod)
	req.Clean()

	rows, total, err := models.RoleFecthList(assoc, req.RoleFilter, req.Sorter, req.Page, scope)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
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

// Restore 恢复角色
// @Summary 恢复角色
// @Description 恢复角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body RestoreReq true "请求体"
// @Success 200 {object} RestoreRes
// @Router /role/restore [put]
func Restore(c *gin.Context) {
	req := RestoreReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(req)
	row, err := models.RoleRestore(req.IDList)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "数据恢复失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": row,
		"msg":  "success",
	})
}

// Edit 编辑角色
// @Summary 编辑角色
// @Description 编辑角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param data body EditBodyReq true "请求体"
// @Success 200 {object} EditRes
// @Router /role/edit/{id} [put]
func Edit(c *gin.Context) {
	uri := EditReq{}
	if err := c.ShouldBindUri(&uri); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(uri)
	body := EditBodyReq{ID: uri.ID}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(body)
	fieldMap := body.ToFiledMap()
	log.Println(fieldMap)
	rows, err := models.RoleEdit(fieldMap, uri.ID)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "数据更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rows,
		"msg":  "success",
	})
}

func EditBatchEnable(c *gin.Context) {
	req := EditBatchEnableReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(req)
	rows, err := models.RoleEditBatchEnable(req.IDList, req.Enabled)
	if err != nil {
		utils.Logger().Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "数据更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rows,
		"msg":  "success",
	})
}
