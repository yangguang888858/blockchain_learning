package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ctxKey string

const (
	ctxKeyId ctxKey = "id"
)

type AuditFields struct {
	CreateBy string `gorm:"size:100;comment:创建人" json:"createBy"`
	UpdateBy string `gorm:"size:100;comment:更新人" json:"updateBy"`
	DeleteBy string `gorm:"size:100;comment:删除人" json:"deleteBy"`
}

// 用户结构体
type User struct {
	ID         uint64         `gorm:"primaryKey;comment:主键" json:"id"`
	Name       string         `gorm:"size:100;comment:姓名" json:"name"`
	Posts      []Post         `json:"posts"`
	PostCount  uint64         `gorm:"default:0;comment:文章数量" json:"postCount"`
	CreateTime time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createTime"`
	UpdateTime time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updateTime"`
	DeteleTime gorm.DeletedAt `gorm:"comment:删除时间" json:"deleteTime"`
	Audit      AuditFields    `gorm:"embedded" json:"audit"`
}

// 文章结构体
type Post struct {
	ID            uint64         `gorm:"primaryKey;comment:主键" json:"id"`
	UserID        uint64         `gorm:"comment:用户id" json:"userID"`
	User          User           `json:"user"`
	Comments      []Comment      `json:"comments"`
	CommentCount  uint64         `gorm:"default:0;comment:评论数量" json:"commentCount"`
	CommentStatus string         `gorm:"size:100;comment:评论状态" json:"commentStatus"`
	Title         string         `gorm:"size:100;comment:标题" json:"title"`
	Content       string         `gorm:"comment:内容" json:"content"`
	CreateTime    time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createTime"`
	UpdateTime    time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updateTime"`
	DeteleTime    gorm.DeletedAt `gorm:"comment:删除时间" json:"deleteTime"`
	Audit         AuditFields    `gorm:"embedded" json:"audit"`
}

// 评论结构体
type Comment struct {
	ID         uint64         `gorm:"primaryKey;comment:主键" json:"id"`
	UserID     uint64         `gorm:"comment:用户id" json:"userID"`
	User       User           `json:"user"`
	PostID     uint64         `gorm:"comment:文章id" json:"postID"`
	Post       Post           `json:"post"`
	Content    string         `gorm:"comment:内容" json:"content"`
	CreateTime time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"createTime"`
	UpdateTime time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updateTime"`
	DeteleTime gorm.DeletedAt `gorm:"comment:删除时间" json:"deleteTime"`
	Audit      AuditFields    `gorm:"embedded" json:"audit"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("执行了Post结构体的BeforeCreate()方法....")
	//1.查询该用户有多少篇文章
	var user User
	tx.Where("id", p.UserID).Find(&user)
	postCount := user.PostCount
	//2.更新用户表的post_count字段
	tx.Model(&User{}).Where("id", p.UserID).Update("post_count", postCount+1)
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Println("执行了Comment结构体的AfterDelete()方法....")
	id := GetValue(tx)
	//1.查询该文章有多少评论
	var post Post
	tx.Where("id", id).Preload("Comments").Find(&post)
	commentCount := len(post.Comments)
	if commentCount == 0 {
		//2.更新文章表的comment_status字段
		tx.Model(&Post{}).Where("id", id).Updates(map[string]any{"comment_status": "无评论", "comment_count": 0})
	}
	return nil
}

func WithValue(id uint64) context.Context {
	return context.WithValue(context.Background(), ctxKeyId, id)
}

func GetValue(tx *gorm.DB) uint64 {
	if tx != nil && tx.Statement != nil && tx.Statement.Context != nil {
		if v, ok := tx.Statement.Context.Value(ctxKeyId).(uint64); ok && v != 0 {
			return v
		}
	}
	return 0
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	fmt.Println("执行了Comment结构体的AfterCreate()方法....")
	id := GetValue(tx)
	//1.查询该文章有多少评论
	var post Post
	tx.Where("id", id).Find(&post)
	commentCount := post.CommentCount
	//2.更新文章表的comment_count字段
	tx.Model(&Post{}).Where("id", id).Update("comment_count", commentCount+1)
	return nil
}
