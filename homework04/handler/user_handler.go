package handler

import (
	"errors"
	"homework04/models"
	"homework04/service"
	"homework04/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var createUserRequest *models.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "请求参数校验失败", err.Error())
		return
	}
	service := service.NewService()
	//对密码进行加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserRequest.PassWord), bcrypt.DefaultCost)
	user := &models.User{
		UserName: createUserRequest.UserName,
		PassWord: string(hashedPassword),
		Email:    createUserRequest.Email,
	}
	if user, err := service.CreateUser(user); err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建用户失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "创建用户成功", user, utils.Page{}, "")
	}
}

func Login(c *gin.Context) {
	var loginUserRequest *models.LoginUserRequest
	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "请求参数校验失败", err.Error())
		return
	}
	//根据用户名验证用户是否存在
	service := service.NewService()
	user := models.User{
		UserName: loginUserRequest.UserName,
	}
	if user, err := service.GetUser(&user); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			utils.Error(c, http.StatusBadRequest, "该用户不存在", err)
			return
		}
		utils.Error(c, http.StatusInternalServerError, "服务器内部错误", err)
		return
	} else {
		//判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(loginUserRequest.PassWord)); err != nil {
			utils.Error(c, http.StatusBadRequest, "密码错误", err)
			return
		} else {
			//生成token
			token, err := utils.GenerateToken([]byte("my-secret"), user.ID, user.UserName)
			if err != nil {
				utils.Error(c, http.StatusInternalServerError, "生成token失败", err)
				return
			}
			utils.Success(c, http.StatusOK, "登录成功", user, utils.Page{}, token)
		}
	}

}

func GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	service := service.NewService()
	user := &models.User{
		ID: id,
	}
	if user, err := service.GetUser(user); err != nil {
		utils.Error(c, http.StatusBadRequest, "获取用户信息失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "获取用户信息成功", user, utils.Page{}, "")
	}
}

func UpdateUser(c *gin.Context) {
	var updateUserRequest *models.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		utils.ValidateParams(c, http.StatusUnprocessableEntity, "获取用户修改信息失败", err.Error())
		return
	}
	service := service.NewService()
	hashedPassWord, _ := bcrypt.GenerateFromPassword([]byte(updateUserRequest.PassWord), bcrypt.DefaultCost)
	updateUser := &models.User{
		ID:       updateUserRequest.ID,
		UserName: updateUserRequest.UserName,
		PassWord: string(hashedPassWord),
		Email:    updateUserRequest.Email,
	}
	if user, err := service.UpdateUser(updateUser); err != nil {
		utils.Error(c, http.StatusInternalServerError, "修改用户信息失败", err)
		return
	} else {
		utils.Success(c, http.StatusOK, "修改用户信息成功", user, utils.Page{}, "")
	}
}
