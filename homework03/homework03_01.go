package main

import (
	"fmt"
	"gin/models"
	"gin/util"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

func create() {

	db := util.NewTestDB("homework03")
	//创建表
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		fmt.Printf("创建表失败%v\n", err)
	}

	//张三
	//1.新增用户
	user1 := models.User{
		Name: "张三",
	}
	if err := db.Create(&user1).Error; err != nil {
		fmt.Printf("新增用户失败%v\n", err)
	}
	//2.新增文章
	posts1 := []models.Post{
		{
			User:          user1,
			Title:         "张三的文章1--标题",
			Content:       "张三的文章1--内容",
			CommentStatus: "正常",
		},
		{
			User:          user1,
			Title:         "张三的文章2--标题",
			Content:       "张三的文章2--内容",
			CommentStatus: "正常",
		},
	}
	if err := db.Create(&posts1).Error; err != nil {
		fmt.Printf("新增文章失败%v\n", err)
	}
	//2.新增评论
	comments11 := []models.Comment{
		{User: user1, Post: posts1[0], Content: "张三的文章1--评论1"},
		{User: user1, Post: posts1[0], Content: "张三的文章1--评论2"},
	}
	createCtx11 := models.WithValue(posts1[0].ID)
	if err := db.WithContext(createCtx11).Create(&comments11).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}
	comments12 := []models.Comment{
		{User: user1, Post: posts1[1], Content: "张三的文章2--评论1"},
		{User: user1, Post: posts1[1], Content: "张三的文章2--评论2"},
	}
	createCtx12 := models.WithValue(posts1[1].ID)
	if err := db.WithContext(createCtx12).Create(&comments12).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}

	//李四
	//1.新增用户
	user2 := models.User{
		Name: "李四",
	}
	if err := db.Create(&user2).Error; err != nil {
		fmt.Printf("新增用户失败%v\n", err)
	}
	//2.新增文章
	posts2 := []models.Post{
		{
			User:          user2,
			Title:         "李四的文章1--标题",
			Content:       "李四的文章1--内容",
			CommentStatus: "正常",
		},
		{
			User:          user2,
			Title:         "李四的文章2--标题",
			Content:       "李四的文章2--内容",
			CommentStatus: "正常",
		},
		{
			User:          user2,
			Title:         "李四的文章3--标题",
			Content:       "李四的文章3--内容",
			CommentStatus: "正常",
		},
	}
	if err := db.Create(&posts2).Error; err != nil {
		fmt.Printf("新增文章失败%v\n", err)
	}
	//2.新增评论
	comments21 := []models.Comment{
		{User: user2, Post: posts2[0], Content: "李四的文章1--评论1"},
		{User: user2, Post: posts2[0], Content: "李四的文章1--评论2"},
		{User: user2, Post: posts2[0], Content: "李四的文章1--评论3"},
	}
	createCtx21 := models.WithValue(posts2[0].ID)
	if err := db.WithContext(createCtx21).Create(&comments21).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}
	comments22 := []models.Comment{
		{User: user2, Post: posts2[1], Content: "李四的文章2--评论1"},
		{User: user2, Post: posts2[1], Content: "李四的文章2--评论2"},
		{User: user2, Post: posts2[1], Content: "李四的文章2--评论3"},
	}
	createCtx22 := models.WithValue(posts2[1].ID)
	if err := db.WithContext(createCtx22).Create(&comments22).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}
	comments23 := []models.Comment{
		{User: user2, Post: posts2[2], Content: "李四的文章3--评论1"},
		{User: user2, Post: posts2[2], Content: "李四的文章3--评论2"},
		{User: user2, Post: posts2[2], Content: "李四的文章3--评论3"},
		{User: user2, Post: posts2[2], Content: "李四的文章3--评论4"},
	}
	createCtx23 := models.WithValue(posts2[2].ID)
	if err := db.WithContext(createCtx23).Create(&comments23).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}
}
