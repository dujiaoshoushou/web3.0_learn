package comment

import "task-week-four/models"

type AddCommentReq struct {
	models.Comment
	Status  int     `json:"status" binding:"required"`
	Content string  `json:"content" binding:"required"`
	PostID  *uint64 `json:"post_id" binding:"required,gt=0"`
}

func (req *AddCommentReq) ToPost() *models.Comment {
	row := req.Comment
	row.PostID = req.PostID
	row.Content = req.Content
	row.UserID = req.UserID
	return &row
}

type GetListReq struct {
	models.CommentFilter
	models.Sorter
	models.Page
}

func (req *GetListReq) Clean() {
	req.CommentFilter.Clean()
	req.Sorter.Clean()
	req.Page.Clean()
}
