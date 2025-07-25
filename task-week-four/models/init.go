package models

import (
	"errors"
	"log"
	"log/slog"
	"task-week-four/utils"
)

func Init() {
	// 迁移数据库
	migrate()
	// 初始化数据
	//seed()
}

func migrate() {
	if err := utils.DB().AutoMigrate(&Role{}, &User{}, &Post{}, &Comment{}); err != nil {
		slog.Error("AutoMigrate err:", err)
		log.Fatalln(errors.New("AutoMigrate err"), err)
	}
}

func seed() {
	roleSeeder()
}
