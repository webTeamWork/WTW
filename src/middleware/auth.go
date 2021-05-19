package middleware

import (
	"forum/src/model"
	"forum/src/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiUnauth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": -1,
		"msg":  msg,
	})
	c.Abort()
}

func UserAuth(c *gin.Context) {
	token := c.GetHeader("User-Token")
	claims, err := jwt.ParseJWT(token)
	if err != nil {
		apiUnauth(c, err.Error())
		return
	}

	var userType int8
	err = model.DB.Get(&userType, "SELECT user_type FROM user WHERE user_id = ?", claims.UserID)
	if err != nil || userType != model.UserTypeUser {
		apiUnauth(c, "无权限访问该接口")
		return
	}

	c.Set("UserID", claims.UserID)
	c.Next()
}
