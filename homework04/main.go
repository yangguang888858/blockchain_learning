package main

import (
	"fmt"
	"homework04/config"
	"homework04/handler"
	"homework04/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.UseRawPath = true          //设置使用原始路径
	r.UnescapePathValues = false //设置不让路径参数自动解码
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "health",
			"message": "ok",
		})
	})
	//公开的接口
	//注册
	r.POST("/register", handler.Register)
	//登录
	r.POST("/login", handler.Login)
	//受保护的接口
	protect := r.Group("")
	protect.Use(middleware.Auth([]byte("my-secret")))
	{
		//----用户模块----
		//获取用户
		protect.GET("/user/:id", handler.GetUser)
		//修改用户信息
		protect.PUT("/user", handler.UpdateUser)

		//----文章模块----
		//创建文章
		protect.POST("/post", handler.CreatePost)
		//查询文章列表
		protect.GET("/post/title/:title/current/:current/size/:size", handler.GetPostList)
		//查询文章详情
		protect.GET("/post/:id", handler.GetPostDetail)
		//修改文章
		protect.PUT("/post", handler.UpdatePost)
		//删除文章
		protect.DELETE("/post", handler.DeletePost)

		//----评论模块----
		//创建评论
		protect.POST("/comment", handler.CreateComment)
		//根据postID查询评论
		protect.GET("/comment/:postID", handler.GetCommentsByPostID)

	}
	//读取配置
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("读取配置失败", err)
		return
	}
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	r.Run(addr)
}
