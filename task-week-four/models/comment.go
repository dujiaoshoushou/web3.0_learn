package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"strings"
	"task-week-four/utils"
)

type Comment struct {
	Model
	Content string  `gorm:"type:varchar(256);not null" json:"content"`
	UserID  *uint64 `gorm:"type:bigint;not null" json:"user_id"`
	Status  int     `gorm:"type:tinyint;not null;default:1;" json:"status"`
	User    User    `gorm:"foreignKey:UserID;references:;" json:"user"`
	PostID  *uint64 `gorm:"type:bigint;not null" json:"post_id"`
	Post    Post    `gorm:"foreignKey:PostID;references:;" json:"post"`
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
func CommentInsert(row *Comment) (err error) {
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
type CommentFilter struct {
	Content *string `form:"content" binding:"omitempty"`
	PostID  *uint64 `form:"post_id" binding:"omitempty"`
}

func (f *CommentFilter) Clean() {
	if f.Content == nil {
		temp := ""
		f.Content = &temp
	}
	if f.PostID == nil {
		temp := uint64(0)
		f.PostID = &temp
	}
}

func CommentFecthList(assoc bool, filter CommentFilter, sorter Sorter, page Page, scope uint8) ([]*Comment, int64, error) {
	query := utils.DB().Model(&Comment{})

	switch scope {
	case SCOPE_ALL:
		query.Unscoped()
	case SCOPE_DELETED:
		query.Unscoped().Where("deleted_at is not null")
	case SCOPE_UNDELETED:
		fallthrough
	default:
	}

	if *filter.Content != "" {
		query = query.Where(" content like  ?", "%"+*filter.Content+"%")
	}
	if *filter.PostID != uint64(0) {
		query = query.Where("post_id = ?", *filter.PostID)
	}
	// 查询总数
	total := int64(0)
	if err := query.Count(&total).Error; err != nil {
		slog.Error("CommentFecthList error", "error", err)
	}
	// 排序
	query.Order(fmt.Sprintf("`%s` %s", *sorter.SortField, strings.ToUpper(*sorter.SortMethod)))

	// 分页
	var rows []*Comment
	if *page.PageSize > 0 {
		offset := (*page.PageNum - 1) * *page.PageSize
		query = query.Offset(offset).Limit(*page.PageSize)
	}

	if err := query.Find(&rows).Error; err != nil {
		slog.Error("CommentFecthList error", "error", err)
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
// - *Comment: 查询到的文章。
// - error: 如果查询过程中发生错误，返回相应的错误信息；否则返回 nil。
func CommentFetchRow(assoc bool, where any, args ...any) (*Comment, error) {
	db := utils.DB()
	row := &Comment{}
	if assoc {
		db = db.Preload("Post")
	}
	if err := db.Where(where, args...).First(row).Error; err != nil {
		slog.Error("CommentFetchRow error", "error", err)
		return nil, err
	}
	return row, nil
}
