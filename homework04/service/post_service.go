package service

import (
	"fmt"
	"homework04/models"
	"homework04/utils"
)

// 创建文章
func (service *Service) CreatePost(post *models.Post) (*models.Post, error) {
	db, _ := utils.Connect()
	params := make(map[string]any, 0)
	params["createBy"] = post.UserName
	params["updateBy"] = post.UserName
	ctx := models.SetCurrentOperator(params)
	if err := db.WithContext(ctx).Create(post).Error; err != nil {
		fmt.Println("创建文章失败", err)
		return nil, err
	}
	//查询文章并返回
	if err := db.Where("id", post.ID).Find(&post).Error; err != nil {
		fmt.Printf("查询文章:%v失败\n", post.ID)
		return nil, err
	}
	return post, nil
}

// 查询文章列表
func (service *Service) GetPostList(getPostRequest *models.GetPostRequest) ([]*models.Post, error) {
	db, _ := utils.Connect()
	current := getPostRequest.Current
	size := getPostRequest.Size
	posts := make([]*models.Post, 0)
	if getPostRequest.Title != "" {
		if err := db.Model(&models.Post{}).Where("title like ?", "%"+getPostRequest.Title+"%").Offset((int(getPostRequest.Current) - 1) * int(getPostRequest.Size)).Limit(int(getPostRequest.Size)).Order("create_at desc").Find(&posts).Error; err != nil {
			fmt.Println("查询文章列表失败", err)
			return nil, err
		}
	} else {
		if err := db.Model(&models.Post{}).Scopes(Paginate(current, size)).Order("create_at desc").Find(&posts).Error; err != nil {
			fmt.Println("查询文章列表失败", err)
			return nil, err
		}
	}
	return posts, nil
}

// 查询文章列表条数
func (service *Service) GetPostListCount(getPostRequest *models.GetPostRequest) (int64, error) {
	db, _ := utils.Connect()
	posts := make([]*models.Post, 0)
	var count int64
	if getPostRequest.Title != "" {
		if err := db.Model(&models.Post{}).Where("title like ?", "%"+getPostRequest.Title+"%").Find(&posts).Count(&count).Error; err != nil {
			fmt.Println("查询文章列表条数失败", err)
			return 0, err
		}
	} else {
		if err := db.Model(&models.Post{}).Find(&posts).Count(&count).Error; err != nil {
			fmt.Println("查询文章列表条数失败", err)
			return 0, err
		}
	}
	return count, nil
}

// 查询文章详情
func (service *Service) GetPostDetail(requestPost *models.Post) (*models.Post, error) {
	db, _ := utils.Connect()
	var post *models.Post
	if requestPost.ID != 0 {
		if err := db.Where("id", requestPost.ID).First(&post).Error; err != nil {
			fmt.Println("查询文章详情失败", err)
			return nil, err
		}
	} else if requestPost.Title != "" {
		if err := db.Where("title like ? ", "%"+requestPost.Title+"%").First(&post).Error; err != nil {
			return nil, err
		}
	}
	return post, nil
}

// 修改文章
func (service *Service) UpdatePost(updatePost *models.Post) (*models.Post, error) {
	db, _ := utils.Connect()
	params := make(map[string]any, 0)
	params["updateBy"] = updatePost.UserName
	ctx := models.SetCurrentOperator(params)
	if err := db.WithContext(ctx).Model(&models.Post{}).Where("id", updatePost.ID).
		Updates(map[string]any{"title": updatePost.Title, "content": updatePost.Content}).Error; err != nil {
		fmt.Println("修改文章失败", err)
		return nil, err
	}
	db.Where("id", updatePost.ID).Find(&models.Post{})
	return updatePost, nil
}

// 删除文章
func (service *Service) DeletePost(deletePost *models.Post) (*models.Post, error) {
	db, _ := utils.Connect()
	params := make(map[string]any, 0)
	params["deleteBy"] = deletePost.UserName
	params["operateType"] = "delete"
	ctx := models.SetCurrentOperator(params)
	//先更新删除人
	if err := db.WithContext(ctx).
		Model(&models.Post{}).
		Where("id", deletePost.ID).
		Update("delete_by", deletePost.UserName).Error; err != nil {
		fmt.Println("更新删除人失败", err)
		return nil, err
	}
	//再删除文章
	if err := db.WithContext(ctx).
		Model(&models.Post{}).
		Delete(deletePost).Error; err != nil {
		fmt.Println("删除文章失败", err)
		return nil, err
	}
	db.Where("id", deletePost.ID).Find(&models.Post{})
	return deletePost, nil
}
