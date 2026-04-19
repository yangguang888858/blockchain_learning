package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement;comment:主键" json:"id" `
	UserID   uint64 `gorm:"not null;comment:用户id" json:"userID" `
	User     User
	PostID   uint64 `gorm:"not null;comment:文章id" json:"postID" `
	Post     Post
	Content  string         `gorm:"size:100;not null;comment:评论内容" json:"content" `
	CreateAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createAt"`
	UpdateAt time.Time      `gorm:"autoUpdateTime;comment:修改时间" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"comment:创建时间" json:"deleteAt"`
	Audit    AuditFields    `gorm:"embedded"`
}

type CreateCommentRequest struct {
	PostID  uint64 `json:"postID" binding:"required"`
	Content string `json:"content" binding:"required,max=100"`
}
