package main

import (
	"fmt"
	"gin/models"
	"gin/util"
)

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func main() {

	db := util.NewTestDB("homework03")
	// //创建表
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		fmt.Printf("创建表失败%v\n", err)
	}

	//1.新增用户
	user3 := models.User{
		Name: "王五",
	}
	if err := db.Create(&user3).Error; err != nil {
		fmt.Printf("新增用户失败%v\n", err)
	}
	//2.新增文章
	posts3 := []models.Post{
		{
			User:          user3,
			Title:         "王五的文章1--标题",
			Content:       "王五的文章1--内容",
			CommentStatus: "正常",
		},
		{
			User:          user3,
			Title:         "王五的文章2--标题",
			Content:       "王五的文章2--内容",
			CommentStatus: "正常",
		},
	}
	if err := db.Create(&posts3).Error; err != nil {
		fmt.Printf("新增文章失败%v\n", err)
	}
	//3.新增评论
	comments31 := []models.Comment{
		{User: user3, Post: posts3[0], Content: "王五的文章1--评论1"},
		{User: user3, Post: posts3[0], Content: "王五的文章1--评论2"},
	}
	createCtx31 := models.WithValue(posts3[0].ID)
	if err := db.WithContext(createCtx31).Create(&comments31).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}

	comments32 := []models.Comment{
		{User: user3, Post: posts3[1], Content: "王五的文章2--评论1"},
		{User: user3, Post: posts3[1], Content: "王五的文章2--评论2"},
		{User: user3, Post: posts3[1], Content: "王五的文章2--评论3"},
	}
	createCtx32 := models.WithValue(posts3[1].ID)
	if err := db.WithContext(createCtx32).Create(&comments32).Error; err != nil {
		fmt.Printf("新增评论失败%v\n", err)
	}

	//删除
	var post models.Post
	db.Where("post_id", 1).Preload("Comments").Find(&post)
	fmt.Printf("删除前<<<id:%v,标题:%v,内容:%v,评论数量:%v,状态:%v\n",
		post.ID, post.Title, post.Content, post.CommentCount, post.CommentStatus)
	deleteCtx := models.WithValue(post.ID)
	db.Unscoped().WithContext(deleteCtx).Where("post_id", 1).Delete(&models.Comment{})
	db.Where("post_id", 1).Preload("Comments").Find(&post)
	fmt.Printf("删除后>>>id:%v,标题:%v,内容:%v,评论数量:%v,状态:%v\n",
		post.ID, post.Title, post.Content, post.CommentCount, post.CommentStatus)

}
