package role

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"task-week-four/models"
	"task-week-four/utils"
)

// GetRowReq 获取单行数据请求参数
type GetRowReq struct {
	ID uint `form:"id" binding:"required,gt=0"`
}

type GetListReq struct {
	models.RoleFilter
	models.Sorter
	models.Page
}

func (req *GetListReq) Clean() {
	req.RoleFilter.Clean()
	req.Sorter.Clean()
	req.Page.Clean()
}

type AddReq struct {
	models.Role
	Title string `json:"title" binding:"required,roleTitleUnique"`
	Key   string `json:"key" binding:"required"`
}

// ToRole 将AddReq转换为Role
func (req AddReq) ToRole() *models.Role {
	row := &req.Role
	row.Key = req.Key
	row.Title = req.Title
	return row
}

type DeleteReq struct {
	IDList []uint `json:"id_list" binding:"required,gt=0"`
	Force  bool   `json:"force" default:"false",binding:""`
}

type RestoreReq struct {
	IDList []uint `json:"id_list" binding:"required,gt=0"`
}

type EditReq struct {
	ID uint `uri:"id" binding:"required,gt=0"`
}

type EditBodyReq struct {
	ID      uint
	Title   *string `json:"title" field:"title" binding:"omitempty,roleTitleUnique"`
	Key     *string `json:"key" field:"key"`
	Enabled *bool   `json:"enabled" field:"enabled"`
	Weight  *int    `json:"weight" field:"weight"`
	Comment *string `json:"comment" field:"comment"`
}

func (req EditBodyReq) ToFiledMap() models.FieldMap {
	m := models.FieldMap{}

	reqType := reflect.TypeOf(req)
	reqValue := reflect.ValueOf(req)
	for i, nums := 0, reqType.NumField(); i < nums; i++ {
		fieldTag := reqType.Field(i).Tag.Get("field")
		if fieldTag == "" {
			continue
		}
		if !reqValue.Field(i).IsNil() {
			if fieldTag == "some_filed" {
				// 特殊情况处理
			} else {
				m[fieldTag] = reqValue.Field(i).Elem().Interface()
			}
		}

	}
	return m
}

type EditBatchEnableReq struct {
	IDList  []uint `json:"id_list" binding:"required,gt=0"`
	Enabled bool   `json:"enabled"`
}

func roleTitleUnique(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().Interface().(string)
	id := fieldLevel.Parent().FieldByName("ID").Interface().(uint64)
	row := models.Role{}
	utils.DB().Where("title = ? && id != ?", value, id).Unscoped().First(&row)
	return row.ID == 0
}
