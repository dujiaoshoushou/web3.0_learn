package gorm

import (
	"gorm.io/gorm"
	"log"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	gorm.Model
	UserName  string
	Age       int
	Password  string
	Email     string
	Posts     []Post    `gorm:"foreignkey:UserID;references:;"`
	Comments  []Comment `gorm:"foreignkey:UserID;references:;"`
	PostCount int
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	User          User `gorm:"foreignkey:UserID;references:;"`
	UserID        *uint
	Comments      []Comment `gorm:"foreignkey:PostID;references:;"`
	CommentCount  int
	CommentStatus string
}

type Comment struct {
	gorm.Model
	Content string
	UserID  *uint
	PostID  *uint
	Post    Post `gorm:"foreignkey:PostID;references:;"`
	User    User `gorm:"foreignkey:UserID;references:;"`
}

//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	log.Println("userID:", p.UserID)
	if err := tx.Model(&user).First(&user, "id = ?", p.UserID).Error; err != nil {
		log.Fatalln(err)
		return err
	}
	user.PostCount++
	if err := tx.Save(&user).Error; err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	if err := tx.Model(&post).First(&post, "id = ?", c.PostID).Error; err != nil {
		log.Fatalln(err)
		return err
	}
	post.CommentCount--
	if post.CommentCount <= 0 {
		post.CommentStatus = "无评论"
	}
	if err := tx.Save(&post).Error; err != nil {
		log.Fatalln(err)
		return err
	}
	return nil

}
