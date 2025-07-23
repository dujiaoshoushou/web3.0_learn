package gorm

import (
	"testing"
)

func TestInitAutoMigrate(t *testing.T) {
	InitAutoMigrate()
}

func TestInitData(t *testing.T) {
	initData()
}

func TestGetPostAndComment(t *testing.T) {
	GetPostAndComment("王五")
}

func TestGetMaxCommentsPost(t *testing.T) {
	GetMaxCommentsPost()
}

func TestInsertPost(t *testing.T) {
	InsertPost()
}

func TestDeleteComments(t *testing.T) {
	DeleteComments()
}
