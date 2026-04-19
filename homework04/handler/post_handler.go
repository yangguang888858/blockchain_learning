package handler

import (
	"errors"
	"fmt"
	"homework04/models"
	"homework04/service"
	"homework04/utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	var createPostRequest models.CreatePostRequest
	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "获取创建文章信息失败", err.Error())
		return
	}
	//从上下文中获取用户信息
	userID, okUserID := c.Get("userID")
	userName, okUserName := c.Get("userName")
	if !okUserID || !okUserName {
		utils.Error(c, http.StatusInternalServerError, "获取用户信息失败", "")
		return
	}
	post := &models.Post{
		UserID:   userID.(uint64),
		UserName: userName.(string),
		Title:    createPostRequest.Title,
		Content:  createPostRequest.Content,
	}
	service := service.NewService()
	if post, err := service.CreatePost(post); err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建文章失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "创建文章信息成功", post, utils.Page{}, "")
	}
}

func GetPostList(c *gin.Context) {
	title := c.Param("title")
	currentStr := c.Param("current")
	sizeStr := c.Param("size")
	current, _ := strconv.ParseUint(currentStr, 10, 64)
	size, _ := strconv.ParseUint(sizeStr, 10, 64)
	//解码title参数
	fmt.Println("解码前:", title)
	title, err := url.QueryUnescape(title)
	if err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "解析title参数失败", err.Error())
		return
	}
	var getPostRequest models.GetPostRequest
	getPostRequest.Title = title
	getPostRequest.Current = current
	getPostRequest.Size = size
	fmt.Println("解码后:", getPostRequest.Title)
	service := service.NewService()
	count, err := service.GetPostListCount(&getPostRequest)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "查询文章列表条数失败", err)
		return
	}
	//计算分页参数
	var pages uint64
	if uint64(count)%getPostRequest.Size == 0 {
		pages = uint64(count) / getPostRequest.Size
	} else {
		pages = uint64(count)/getPostRequest.Size + 1
	}
	page := utils.Page{
		Current: getPostRequest.Current,
		Size:    getPostRequest.Size,
		Total:   uint64(count),
		Pages:   pages,
	}
	if posts, err := service.GetPostList(&getPostRequest); err != nil {
		utils.Error(c, http.StatusBadRequest, "查询文章列表失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "查询文章列表成功", posts, page, "")
	}
}

func GetPostDetail(c *gin.Context) {
	service := service.NewService()
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	post := &models.Post{
		ID: id,
	}
	if posts, err := service.GetPostDetail(post); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			utils.Error(c, http.StatusBadRequest, "此文章不存在", err)
			return
		}
		utils.Error(c, http.StatusBadRequest, "查询文章详情失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "查询文章详情成功", posts, utils.Page{}, "")
	}
}

func UpdatePost(c *gin.Context) {
	var updatePostRequest *models.UpdatePostRequest
	if err := c.ShouldBindJSON(&updatePostRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "获取文章修改信息失败", err.Error())
		return
	}
	service := service.NewService()
	//先查询文章
	getPost := &models.Post{
		ID: updatePostRequest.ID,
	}
	post, err := service.GetPostDetail(getPost)
	if errors.Is(gorm.ErrRecordNotFound, err) || post.ID == 0 {
		utils.Error(c, http.StatusBadRequest, "此文章不存在", "")
		return
	}
	//从上下文中获取用户信息
	userID, okUserID := c.Get("userID")
	userName, okUserName := c.Get("userName")
	if !okUserID || !okUserName {
		utils.Error(c, http.StatusInternalServerError, "获取用户信息失败", "")
		return
	}
	//将要修改文章的userID和token中的userID进行匹配
	//如果匹配不上,说明修改的是别人的文章,提示权限不足
	if post.UserID != userID.(uint64) {
		utils.Error(c, http.StatusUnauthorized, "只能修改自己的文章", "")
		return
	}
	//再修改文章
	updatePost := &models.Post{
		ID:       updatePostRequest.ID,
		UserName: userName.(string),
		Title:    updatePostRequest.Title,
		Content:  updatePostRequest.Content,
	}
	if post, err := service.UpdatePost(updatePost); err != nil {
		utils.Error(c, http.StatusInternalServerError, "修改文章失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "修改文章成功", post, utils.Page{}, "")
	}
}

func DeletePost(c *gin.Context) {
	var deletePostRequest *models.DeletePostRequest
	if err := c.ShouldBindJSON(&deletePostRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "获取文章删除信息失败", err.Error())
		return
	}
	service := service.NewService()
	//先查询文章
	getPost := &models.Post{
		ID: deletePostRequest.ID,
	}
	post, err := service.GetPostDetail(getPost)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) || post.ID == 0 {
			utils.Error(c, http.StatusBadRequest, "此文章不存在", "")
		}
		return
	}
	//从上下文中获取用户信息
	userID, okUserID := c.Get("userID")
	userName, okUserName := c.Get("userName")
	if !okUserID || !okUserName {
		utils.Error(c, http.StatusInternalServerError, "获取用户信息失败", "")
		return
	}
	//将要删除文章的userID和token中的userID进行匹配
	//如果匹配不上,说明修改的是别人的文章,提示权限不足
	if post.UserID != userID.(uint64) {
		utils.Error(c, http.StatusUnauthorized, "只能删除自己的文章", "")
		return
	}
	//再删除文章
	deletePost := &models.Post{
		ID:       deletePostRequest.ID,
		UserName: userName.(string),
	}
	if post, err := service.DeletePost(deletePost); err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除文章失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "删除文章成功", post, utils.Page{}, "")
	}
}
