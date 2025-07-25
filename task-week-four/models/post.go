package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"strings"
	"task-week-four/utils"
)

type Post struct {
	Model
	Title    string    `gorm:"type:varchar(128);not null" json:"title"`
	Content  string    `gorm:"type:varchar(256);not null" json:"content"`
	UserID   *uint64   `gorm:"type:bigint;not null" json:"user_id"`
	Status   int       `gorm:"type:tinyint;not null;default:1;" json:"status" default:"1"`
	User     User      `gorm:"foreignKey:UserID;references:;" json:"user"`
	Comments []Comment `gorm:"foreignKey:PostID;references:;" json:"comment"`
}

// PostInsert 方法用于将 Post 结构体插入到数据库中。
// 它使用了 GORM 的事务功能，确保插入操作的原子性。
// 如果插入过程中发生错误，将回滚整个事务。
// 如果插入成功，将提交事务。
// PostInsert 方法用于将 Post 结构体插入到数据库中。
// 它使用了 GORM 的事务功能，确保插入操作的原子性。
// 如果插入过程中发生错误，将回滚整个事务。
// 如果插入成功，将提交事务。
//
// 参数:
// - row: 要插入的 Post 结构体指针。
//
// 返回值:
// - error: 如果插入过程中发生错误，返回相应的错误信息；否则返回 nil。
//
// 注意:
// - 该方法使用了 GORM 的 Transaction 方法，确保插入操作的原子性。
// - 如果插入过程中发生错误，将回滚整个事务。
// - 如果插入成功，将提交事务。
// - 插入操作将使用默认的数据库连接。
func PostInsert(row *Post) (err error) {
	return utils.DB().Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(row).Error; err != nil {
			utils.Logger().Error("保存文章失败", "error", err, "row", row)
			return err
		}
		return nil
	})
}

// PostFilter 结构体用于定义查询文章时的过滤条件。
// 它包含一个 Title 字段，用于过滤文章标题。
//
// 字段:
// - Title: 文章标题的过滤条件，类型为 *string。
type PostFilter struct {
	Title *string `form:"title" binding:"omitempty"`
}

// clean 方法用于处理 Filter 结构体中的字段，确保它们不为 nil。
// 如果 Title 字段为 nil，它将被初始化为一个空字符串。
// 这是为了避免在后续的操作中出现 nil 指针异常。
func (f *PostFilter) Clean() {
	if f.Title == nil {
		temp := ""
		f.Title = &temp
	}
}

// PostFecthList 方法用于从数据库中查询文章列表。
// 它接受多个参数，用于指定查询条件、排序方式、分页信息和查询范围。
//
// 参数:
// - assoc: 是否关联查询相关数据，类型为 bool。
// - filter: 查询条件，类型为 PostFilter。
// - sorter: 排序方式，类型为 Sorter。
// - page: 分页信息，类型为 Page。
// - scope: 查询范围，类型为 uint8。
//
// 返回值:
// - []*Post: 查询到的文章列表。
// - int64: 文章总数。
// - error: 如果查询过程中发生错误，返回相应的错误信息；否则返回 nil。
//
// 注意:
// - 该方法使用了 GORM 的链式操作，通过不同的方法来构建查询条件、排序方式和分页信息。
// - 查询范围可以是 SCOPE_ALL、SCOPE_DELETED 或 SCOPE_UNDELETED。
// - 如果查询范围为 SCOPE_DELETED，将查询已删除的文章。
// - 如果查询范围为 SCOPE_UNDELETED，将查询未删除的文章。
// - 如果查询范围为 SCOPE_ALL，将查询所有文章，包括已删除的文章。
// - 查询结果将按照指定的排序方式进行排序。
func PostFecthList(assoc bool, filter PostFilter, sorter Sorter, page Page, scope uint8) ([]*Post, int64, error) {
	query := utils.DB().Model(&Post{})

	switch scope {
	case SCOPE_ALL:
		query.Unscoped()
	case SCOPE_DELETED:
		query.Unscoped().Where("deleted_at is not null")
	case SCOPE_UNDELETED:
		fallthrough
	default:
	}

	if *filter.Title != "" {
		query = query.Where(" title like  ?", "%"+*filter.Title+"%")
	}

	// 查询总数
	total := int64(0)
	if err := query.Count(&total).Error; err != nil {
		slog.Error("PostFecthList error", "error", err)
	}
	// 排序
	query.Order(fmt.Sprintf("`%s` %s", *sorter.SortField, strings.ToUpper(*sorter.SortMethod)))

	// 分页
	var rows []*Post
	if *page.PageSize > 0 {
		offset := (*page.PageNum - 1) * *page.PageSize
		query = query.Offset(offset).Limit(*page.PageSize)
	}

	if err := query.Find(&rows).Error; err != nil {
		slog.Error("PostFecthList error", "error", err)
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

// PostFetchRow 方法用于从数据库中查询单个文章。
// 它接受多个参数，用于指定查询条件和查询范围。
//
// 参数:
// - assoc: 是否关联查询相关数据，类型为 bool。
// - where: 查询条件，类型为 any。
// - args: 查询条件的参数，类型为 any。
//
// 返回值:
// - *Post: 查询到的文章。
// - error: 如果查询过程中发生错误，返回相应的错误信息；否则返回 nil。
func PostFetchRow(assoc bool, where any, args ...any) (*Post, error) {
	db := utils.DB()
	row := &Post{}
	if assoc {
		db = db.Preload("Post")
	}
	if err := db.Where(where, args...).First(row).Error; err != nil {
		slog.Error("PostFetchRow error", "error", err)
		return nil, err
	}
	return row, nil
}

func PostEdit(fieldMap FieldMap, id uint64) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&Post{}).Where("id = ?", id).Updates(fieldMap)
		err := result.Error
		if err != nil {
			slog.Error("PostEdit error", "error", result.Error)
			rowsNum = 0
			return err
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}

func PostDelete(id uint64, force bool) (int64, error) {
	rowsNum := int64(0)
	err := utils.DB().Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&Post{})
		if force {
			query.Unscoped()
		}
		result := query.Delete(&Post{}, id)
		if result.Error != nil {
			slog.Error("PostDelete error", "error", result.Error)
			return result.Error
		} else {
			rowsNum = result.RowsAffected
		}
		return nil
	})
	return rowsNum, err
}
