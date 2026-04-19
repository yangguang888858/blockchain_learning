package utils

import "github.com/gin-gonic/gin"

type Page struct {
	Current uint64 `json:"current"`
	Size    uint64 `json:"size"`
	Total   uint64 `json:"total"`
	Pages   uint64 `json:"pages"`
}
type Response struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data,omitempty"`
	Page  Page   `json:"page"`
	Token string `json:"token,omitempty"`
	Error any    `json:"error,omitempty"`
}

func Success(c *gin.Context, code int, msg string, data any, page Page, token string) {
	c.JSON(code, Response{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Page:  page,
		Token: token,
	})
}

func Error(c *gin.Context, code int, msg string, err any) {
	c.JSON(code, Response{
		Code:  code,
		Msg:   msg,
		Error: err,
	})
}

func ValidateParams(c *gin.Context, code int, msg string, err string) {
	c.JSON(code, Response{
		Code:  code,
		Msg:   msg,
		Error: err,
	})
}
