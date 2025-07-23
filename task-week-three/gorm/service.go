package gorm

import (
	"fmt"
	"log"
)

//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

func initData() {
	user := User{
		UserName: "张三",
		Password: "123456",
		Age:      20,
		Email:    "zhangsan@example.com",
	}
	if err := DB.Create(&user).Error; err != nil {
		log.Fatalln("创建user", err)
	}
	log.Println(user.ID)
	var p1, p2, p3 Post
	p1.Title = "go"
	p2.Title = "java"
	p3.Title = "linux"
	p1.Content = "go语言从入门到精通"
	p2.Content = "java语言从入门到精通"
	p3.Content = "linux从入门到精通"

	if err := DB.Create([]*Post{&p1, &p2, &p3}).Error; err != nil {
		log.Fatalln("批量创建Post异常", err)
	}

	if err := DB.Model(&user).Association("Posts").Append([]Post{p1, p2, p3}); err != nil {
		log.Fatalln("添加user和post关联关系异常", err)
		return
	}

	var c1, c2, c3 Comment
	c1.Content = "go语言从入门到精通"
	c2.Content = "java语言从入门到精通"
	c3.Content = "linux从入门到精通"
	if err := DB.Create([]*Comment{&c1, &c2, &c3}).Error; err != nil {
		log.Fatalln("批量创建Comment异常", err)
	}
	if err := DB.Model(&p1).Association("Comments").Append([]Comment{c1, c2}); err != nil {
		log.Fatalln("添加post和comments异常", err)
	}
	if err := DB.Model(&p2).Association("Comments").Append([]Comment{c3}); err != nil {
		log.Fatalln("添加post和comments异常", err)
	}
	//if err := DB.Model(&p3).Association("Comments").Append([]Comment{c3, c1, c2}); err != nil {
	//	log.Fatalln("添加post和comments异常", err)
	//}

}

// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func GetPostAndComment(userName string) {
	var user User
	var posts []Post
	comments := []Comment{}
	DB.First(&user, "user_name = ?", userName)
	if err := DB.Model(&user).Association("Posts").Find(&posts); err != nil {
		log.Fatalln(err)
	}
	var postIDs []uint
	for _, post := range posts {
		postIDs = append(postIDs, post.ID)
	}
	if err := DB.Model(&posts).Association("Comments").Find(&comments); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(posts)
	fmt.Println("----------")
	fmt.Println(comments)
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func GetMaxCommentsPost() {

	//comment := Comment{}
	type result struct {
		PostID      uint
		CommentsNum uint
	}

	var posts []Post
	var results []result
	fromSubQuery := DB.Model(&Comment{}).Select("post_id", "count(*) as comments_num ").Group("post_id")

	whereSubQuery := DB.Table("( ? ) as tmp", fromSubQuery).Select("max(tmp.comments_num) as max_comments_num")

	if err := DB.Table("( ? ) as temp", fromSubQuery).Where("temp.comments_num = (?)", whereSubQuery).Find(&results).Error; err != nil {
		log.Fatalln(err)
	}
	var comments []Comment
	for _, res := range results {
		comments = append(comments, Comment{
			PostID: &res.PostID,
		})
	}

	if err := DB.Model(&comments).Association("Post").Find(&posts); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(posts)

}

func InsertPost() {
	var UserID uint = 7

	var p1, p2, p3 Post
	p1.Title = "go"
	p2.Title = "java"
	p3.Title = "linux"
	p1.Content = "go语言从入门到精通"
	p2.Content = "java语言从入门到精通"
	p3.Content = "linux从入门到精通"

	p1.UserID = &UserID
	p2.UserID = &UserID
	p3.UserID = &UserID

	if err := DB.Create([]*Post{&p1, &p2, &p3}).Error; err != nil {
		log.Fatalln("批量创建Post异常", err)
	}

}

func DeleteComments() {
	var c1 Comment
	c1.Content = "go语言从入门到精通"

	if err := DB.Model(&c1).Where("post_id is not null or post_id != 0").First(&c1, "content = ?", &c1.Content).Error; err != nil {
		log.Fatalln(err)
	}
	if err := DB.Model(&c1).Delete(&c1).Error; err != nil {
		log.Fatalln(err)
	}

}
