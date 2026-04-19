package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID       uint64         `gorm:"primaryKey;autoIncrement;comment:主键" json:"id" `
	UserID   uint64         `gorm:"not null;comment:用户id" json:"userID" `
	UserName string         `gorm:"-" json:"userName" `
	User     User           `gorm:"-" json:"user"`
	Title    string         `gorm:"size:100;not null;uniqueKey;comment:文章标题" json:"title" `
	Content  string         `gorm:"size:100;not null;comment:文章内容" json:"content" `
	CreateAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createAt"`
	UpdateAt time.Time      `gorm:"autoUpdateTime;comment:修改时间" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"comment:创建时间" json:"deleteAt"`
	Audit    AuditFields    `gorm:"embedded" json:"audit"`
}

type GetPostRequest struct {
	Title   string `json:"title"`
	Current uint64 `json:"current" binding:"required,min=1"`
	Size    uint64 `json:"size" binding:"required,min=1"`
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required,max=100"`
}

type UpdatePostRequest struct {
	ID      uint64 `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required,max=100"`
}

type DeletePostRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	params := GetCurrentOperator(db)
	post.Audit.CreateBy = params["createBy"].(string)
	post.Audit.UpdateBy = params["updateBy"].(string)
	db.Statement.SetColumn("create_by", params["createBy"].(string))
	db.Statement.SetColumn("update_by", params["updateBy"].(string))
	return nil
}

func (post *Post) BeforeUpdate(db *gorm.DB) error {
	params := GetCurrentOperator(db)
	if updateBy, ok := params["operateType"]; ok && updateBy == "delete" {
		return nil
	}
	post.Audit.UpdateBy = params["updateBy"].(string)
	db.Statement.SetColumn("update_by", params["updateBy"].(string))
	return nil
}

func (post *Post) BeforeDelete(db *gorm.DB) error {
	params := GetCurrentOperator(db)
	post.Audit.DeleteBy = params["deleteBy"].(string)
	db.Statement.SetColumn("delete_by", params["deleteBy"].(string))
	return nil
}
