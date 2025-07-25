package user

import "task-week-four/models"

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,gt=0"`
}

type AddUserReq struct {
	models.User
	UserName string `json:"username" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,gt=0"`
	Email    string `json:"email" binding:"required,gt=0"`
}

func (addUser AddUserReq) ToUser() *models.User {
	row := &addUser.User
	row.UserName = addUser.UserName
	row.Password = addUser.Password
	row.Email = addUser.Email
	return row
}
