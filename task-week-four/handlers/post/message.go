package post

import (
	"reflect"
	"task-week-four/models"
)

type AddPostReq struct {
	models.Post
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (req *AddPostReq) ToPost() *models.Post {
	row := req.Post
	row.Title = req.Title
	row.Content = req.Content
	row.UserID = req.UserID
	return &row
}

type GetRowReq struct {
	Title string `form:"title" binding:"required"`
}

type GetListReq struct {
	models.PostFilter
	models.Sorter
	models.Page
}

func (req *GetListReq) Clean() {
	req.PostFilter.Clean()
	req.Sorter.Clean()
	req.Page.Clean()
}

type EditReq struct {
	ID uint64 `uri:"id" binding:"required,gt=0"`
}

type EditBodyReq struct {
	Title   *string `json:"title" binding:"required" field:"title"`
	Content *string `json:"content" binding:"required" field:"content"`
	Status  *int    `json:"status" binding:"required,oneof=0 1" field:"status"`
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

type DeleteReq struct {
	ID uint64 `uri:"id" binding:"required,gt=0"`
}
