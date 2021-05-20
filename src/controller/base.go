package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiInputErr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "请求参数有误",
	})
	c.Abort()
}

func apiErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  msg,
	})
	c.Abort()
}

func apiOK(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

func bindRequest(c *gin.Context, dst interface{}) error {
	err := c.ShouldBindJSON(dst)
	if err != nil {
		apiInputErr(c)
		return err
	}
	return nil
}
