package middleware

import (
	"homework04/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			utils.Error(c, http.StatusBadRequest, "token为空", "")
			c.Abort() //很关键,不然还会往下走
			return
		}
		tokenSplit := strings.Split(token, " ")
		if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
			utils.Error(c, http.StatusBadRequest, "token格式不正确", "")
			c.Abort()
			return
		}
		token = tokenSplit[1]
		claims, err := utils.ParseToken(token, secret)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "无效的token", "")
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.UserName)
		c.Next()
	}
}
