package service

import (
	"fmt"
	"homework04/models"
	"homework04/utils"
)

// 创建用户方法
func (service *Service) CreateUser(user *models.User) (*models.User, error) {
	db, _ := utils.Connect()
	if err := db.Create(user).Error; err != nil {
		fmt.Println("创建用户失败", err)
		return nil, err
	}
	//查询用户并返回
	if err := db.Where("id", user.ID).Find(&user).Error; err != nil {
		fmt.Printf("查询用户:%v失败\n", user.ID)
		return nil, err
	}
	return user, nil
}

// 查询用户
func (service *Service) GetUser(requestUser *models.User) (*models.User, error) {
	db, _ := utils.Connect()
	var user *models.User
	if requestUser.ID != 0 {
		if err := db.Where("id", requestUser.ID).First(&user).Error; err != nil {
			fmt.Println("获取用户信息失败", err)
			return nil, err
		}
	} else if requestUser.UserName != "" {
		if err := db.Where("user_name", requestUser.UserName).First(&user).Error; err != nil {
			return nil, err
		}
	}
	return user, nil
}

// 修改用户信息
func (service *Service) UpdateUser(updateUser *models.User) (*models.User, error) {
	db, _ := utils.Connect()
	if err := db.Model(&models.User{}).Where("id", updateUser.ID).
		Updates(map[string]any{"user_name": updateUser.UserName, "pass_word": updateUser.PassWord, "email": updateUser.Email}).Error; err != nil {
		fmt.Println("修改用户信息失败", err)
		return nil, err
	}
	db.Where("id", updateUser.ID).Find(&models.User{})
	return updateUser, nil
}
