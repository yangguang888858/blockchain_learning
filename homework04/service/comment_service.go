package service

import (
	"fmt"
	"homework04/models"
	"homework04/utils"
)

// 创建评论
func (service *Service) CreateComment(comment *models.Comment) (*models.Comment, error) {
	db, _ := utils.Connect()
	if err := db.Create(comment).Error; err != nil {
		fmt.Println("创建文章失败", err)
		return nil, err
	}
	//查询评论并返回
	if err := db.Where("id", comment.ID).Find(&comment).Error; err != nil {
		fmt.Printf("查询评论失败:%v失败\n", comment.ID)
		return nil, err
	}
	return comment, nil
}

// 根据postID查询评论
func (service *Service) GetCommentsByPostID(requestComment *models.Comment) ([]*models.Comment, error) {
	db, _ := utils.Connect()
	comments := make([]*models.Comment, 0)
	if requestComment.PostID != 0 {
		if err := db.Where("post_id", requestComment.PostID).Find(&comments).Error; err != nil {
			fmt.Println("根据postID查询评论失败", err)
			return nil, err
		}
	}
	return comments, nil
}
