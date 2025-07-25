package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 通用的查询列表过滤类型
type Filter struct {
	Keyword *string `form:"keyword" binding:"omitempty,gt=0"`
}

type Sorter struct {
	SortField *string `form:"sortField" binding:"omitempty,gt=0"`

	SortMethod *string `form:"sortMethod" binding:"omitempty,oneof=DESC ASC"`
}

type Page struct {
	PageSize *int `form:"pageSize" binding:"omitempty,gt=0"`
	PageNum  *int `form:"pageNum" binding:"omitempty,gt=0"`
}

const (
	PageNumDefault  = 1
	PageSizeDefault = 10
	PageSizeMax     = 100

	SortFiledDefault  = "id"
	SortMethodDefault = "DESC"
)

// clean 方法用于处理 Filter 结构体中的字段，确保它们不为 nil。
// 如果 Keyword 字段为 nil，它将被初始化为一个空字符串。
// 这是为了避免在后续的操作中出现 nil 指针异常。
func (f *Filter) Clean() {
	if f.Keyword == nil {
		temp := ""
		f.Keyword = &temp
	}
}

// clean 方法用于处理 Sorter 结构体中的字段，确保它们不为 nil。
// 如果 SortFiled 字段为 nil，它将被初始化为一个默认值。
// 如果 SortMethod 字段为 nil，它将被初始化为一个默认值。
// 这是为了避免在后续的操作中出现 nil 指针异常。
func (s *Sorter) Clean() {
	if s.SortField == nil {
		temp := SortFiledDefault
		s.SortField = &temp
	}
	if s.SortMethod == nil {
		temp := SortMethodDefault
		s.SortMethod = &temp
	}
}

// clean 方法用于处理 Page 结构体中的字段，确保它们不为 nil。
// 如果 PageNum 字段为 nil，它将被初始化为一个默认值。
// 如果 PageSize 字段为 nil，它将被初始化为一个默认值。
// 如果 PageSize 字段的值大于 PageSizeMax，它将被设置为 PageSizeMax 的值。
// 这是为了避免在后续的操作中出现 nil 指针异常。
func (p *Page) Clean() {
	if p.PageNum == nil {
		temp := PageNumDefault
		p.PageNum = &temp
	}
	if p.PageSize == nil {
		temp := PageSizeDefault
		p.PageSize = &temp
	}
	if *p.PageSize > PageSizeMax {
		temp := PageSizeMax
		p.PageSize = &temp
	}
}

const (
	SCOPE_ALL = iota
	SCOPE_UNDELETED
	SCOPE_DELETED
)

type FieldMap = map[string]any
