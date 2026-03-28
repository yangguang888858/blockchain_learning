package main

import (
	"fmt"
	"gin/models"
	"gin/util"
)

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func query() {
	db := util.NewTestDB("homework03")
	//1.查询某个用户发布的所有文章及其对应的评论信息
	var user models.User
	if err := db.Model(&models.User{}).
		Where("name", "李四").
		Preload("Posts.Comments").
		Find(&user).Error; err != nil {
		fmt.Printf("查询失败%v\n", err)
	}
	for _, post := range user.Posts {
		fmt.Printf("id:%v,标题:%v,内容:%v,评论数量:%v,状态:%v\n",
			post.ID, post.Title, post.Content, post.CommentCount, post.CommentStatus)
		for _, comment := range post.Comments {
			fmt.Printf("id:%v,userId:%v,postId:%v,内容:%v\n",
				comment.ID, comment.UserID, comment.PostID, comment.Content)
		}
		fmt.Println("--------------------------------------")
	}

	//2.查询评论数量最多的文章信息

	//2.1 统计哪篇文章评论数最多
	var data map[string]any
	db.Raw("select posts.id,count(comments.id) commentCount from posts left join comments on posts.id=comments.post_id group by posts.id order by count(comments.id) desc limit 1").Scan(&data)
	fmt.Println("data:", data)

	// 2.2 查询评论数最多的那篇
	var post models.Post
	if data != nil {
		id := data["id"]
		db.Where("id", id).First(&post)
	}
	fmt.Printf("id:%v,标题:%v,内容:%v,评论数量:%v,状态:%v\n",
		post.ID, post.Title, post.Content, post.CommentCount, post.CommentStatus)

}
