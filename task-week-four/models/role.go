package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"strings"
	"task-week-four/utils"
)

type Role struct {
	Model
	Title   string `json:"title" gorm:"uniqueIndex;type:varchar(255)"`
	Key     string `json:"key" gorm:"uniqueIndex;type:varchar(255)"`
	Enabled bool   `json:"enabled" gorm:""`
	Weight  int    `json:"weight" gorm:"index;"`
	Comment string `json:"comment" gorm:"type:text"`
}

func roleSeeder() {
	roles := []Role{
		{Title: "超级管理员", Key: "administor", Enabled: true, Model: Model{ID: 1}},
		{Title: "常规用户", Key: "regular", Enabled: true, Model: Model{ID: 2}},
	}
	for _, row := range roles {
		if err := utils.DB().FirstOrCreate(&row, row.ID).Error; err != nil {
			slog.Error("roleSeeder error", "error", err)
			log.Fatalln("roleSeeder error", err)
		}
	}
}

// 根据条件查询单条角色
// assoc: 是否关联查询
// where: 查询条件
// args: 查询参数
// 返回值:
// role: 角色对象
// err: 错误信息
func RoleFetchRow(assoc bool, where any, args ...any) (*Role, error) {
	db := utils.DB()
	row := &Role{}
	if assoc {
		db = db.Preload("Role")
	}
	if err := db.Where(where, args...).First(row).Error; err != nil {
		slog.Error("RoleFetchRow error", "error", err)
		return nil, err
	}
	return row, nil
}

type RoleFilter struct {
	Keyword *string `form:"keyword" binding:"omitempty,gt=0"`
}

// clean 方法用于处理 Filter 结构体中的字段，确保它们不为 nil。
// 如果 Keyword 字段为 nil，它将被初始化为一个空字符串。
// 这是为了避免在后续的操作中出现 nil 指针异常。
func (f *RoleFilter) Clean() {
	if f.Keyword == nil {
		temp := ""
		f.Keyword = &temp
	}
}

// RoleFecthList 方法用于根据条件查询角色列表。
// assoc: 是否关联查询
// filter: 查询条件
// sorter: 排序条件
// page: 分页条件
// scope: 权限范围
// 返回值:
// rows: 角色列表
// total: 总记录数
// err: 错误信息
func RoleFecthList(assoc bool, filter RoleFilter, sorter Sorter, page Page, scope uint8) ([]*Role, int64, error) {
	query := utils.DB().Model(&Role{})

	switch scope {
	case SCOPE_ALL:
		query.Unscoped()
	case SCOPE_DELETED:
		query.Unscoped().Where("deleted_at is not null")
	case SCOPE_UNDELETED:
		fallthrough
	default:
	}

	if *filter.Keyword != "" {
		query = query.Where(" title like  ?", "%"+*filter.Keyword+"%")
	}

	// 查询总数
	total := int64(0)
	if err := query.Count(&total).Error; err != nil {
		slog.Error("RoleFecthList error", "error", err)
	}
	// 排序
	query.Order(fmt.Sprintf("`%s` %s", *sorter.SortField, strings.ToUpper(*sorter.SortMethod)))

	// 分页
	var rows []*Role
	if *page.PageSize > 0 {
		offset := (*page.PageNum - 1) * *page.PageSize
		query = query.Offset(offset).Limit(*page.PageSize)
	}

	if err := query.Find(&rows).Error; err != nil {
		slog.Error("RoleFecthList error", "error", err)
		return nil, 0, err
	}
	// 是否关联查询
	//if assoc {
	//
	//}
	log.Println("total:", total, "rows:", rows)
	slog.Info("查询row分页结果", "rows", rows, "pageNum", page.PageNum, "pageSize", page.PageSize, "total", total)
	return rows, total, nil
}

func RoleInsert(role *Role) error {
	return utils.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(role).Error; err != nil {
			slog.Error("RoleInsert error", "error", err)
			return err
		}
		return nil
	})
}

func RoleDelete(idList []uint, force bool) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&Role{})
		if force {
			query.Unscoped()
		}
		result := query.Delete(&Role{}, idList)
		if result.Error != nil {
			slog.Error("RoleDelete error", "error", result.Error)
			return result.Error
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}

func RoleRestore(idList []uint) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&Role{}).Unscoped().Where("id in ?", idList).Update("deleted_at", nil)
		if result.Error != nil {
			slog.Error("RoleResotore error", "error", result.Error)
			return result.Error
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}

func RoleEdit(fieldMap FieldMap, id uint) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&Role{}).Where("id = ?", id).Updates(fieldMap)
		err := result.Error
		if err != nil {
			slog.Error("RoleEdit error", "error", result.Error)
			rowsNum = 0
			return err
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}

func RoleEditBatchEnable(idList []uint, enabled bool) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&Role{}).Where("id in ? ", idList).Update("enabled", enabled)
		err := result.Error
		if err != nil {
			slog.Error("RoleEditBatchEnable error", "error", result.Error)
			rowsNum = 0
			return err
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}
