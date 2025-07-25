package models

import (
	"gorm.io/gorm"
	"task-week-four/utils"
)

type User struct {
	Model
	UserName string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Posts    []Post    `gorm:"foreignKey:UserID;references:;"`
	Comments []Comment `gorm:"foreignKey:UserID;references:;"`
}

func UserLogin(username string, password string) (*User, error) {
	user := User{}
	result := utils.DB().Model(&user).Where("user_name = ? and password = ?", username, password).First(&user)
	err := result.Error
	if err != nil {
		utils.Logger().Error("用户登录错误", err)
	}
	return &user, err
}

func UserRegister(row *User) error {

	return utils.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(row).Error; err != nil {
			return err
		}
		return nil
	})
}
