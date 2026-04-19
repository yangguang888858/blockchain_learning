package handler

import (
	"errors"
	"homework04/models"
	"homework04/service"
	"homework04/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context) {
	var createCommentRequest models.CreateCommentRequest
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "获取创建评论信息失败", err.Error())
		return
	}
	//从上下文中获取用户信息
	userID, ok := c.Get("userID")
	if !ok {
		utils.Error(c, http.StatusInternalServerError, "获取用户信息失败", "")
		return
	}
	comment := &models.Comment{
		UserID:  userID.(uint64),
		PostID:  createCommentRequest.PostID,
		Content: createCommentRequest.Content,
	}
	service := service.NewService()
	if comment, err := service.CreateComment(comment); err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建评论失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "创建评论成功", comment, utils.Page{}, "")
	}
}

func GetCommentsByPostID(c *gin.Context) {
	service := service.NewService()
	postIDStr := c.Param("postID")
	postID, _ := strconv.ParseUint(postIDStr, 10, 64)
	comment := &models.Comment{
		PostID: postID,
	}
	if comments, err := service.GetCommentsByPostID(comment); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			utils.Error(c, http.StatusOK, "此文章没有评论", err)
			return
		}
		utils.Error(c, http.StatusBadRequest, "根据postID查询评论失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "根据postID查询评论成功", comments, utils.Page{}, "")
	}
}
