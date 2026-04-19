package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64         `gorm:"primaryKey;autoIncrement;comment:主键" json:"id" `
	UserName string         `gorm:"size:100;not null;uniqueKey;comment:用户名" json:"userName" `
	PassWord string         `gorm:"size:100;not null;comment:密码" json:"passWord" `
	Email    string         `gorm:"size:100;not null;uniqueKey;comment:邮箱" json:"email" `
	CreateAt time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createAt"`
	UpdateAt time.Time      `gorm:"autoUpdateTime;comment:修改时间" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"comment:创建时间" json:"deleteAt"`
	Audit    AuditFields    `gorm:"embedded" json:"audit"`
}

type CreateUserRequest struct {
	UserName string `json:"userName" binding:"required,max=100"`
	PassWord string `json:"passWord" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginUserRequest struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

type UpdateUserRequest struct {
	ID       uint64 `json:"id" binding:"required"`
	UserName string `json:"userName" binding:"max=100"`
	PassWord string `json:"passWord" binding:"min=6"`
	Email    string `json:"email" binding:"email"`
}
