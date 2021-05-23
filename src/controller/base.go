package controller

import (
	"net/http"
	"strconv"

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

func getUserID(c *gin.Context) int {
	return c.GetInt("UserID")
}

func getParamID(c *gin.Context, key string) (int, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		apiErr(c, "接口错误，无法获取ID")
		return 0, false
	} else if id <= 0 {
		apiErr(c, "ID非法")
		return 0, false
	}
	return id, true
}

func getTopicID(c *gin.Context) (int, bool) {
	return getParamID(c, "topic_id")
}

func getCommentID(c *gin.Context) (int, bool) {
	return getParamID(c, "comment_id")
}

func getSectionID(c *gin.Context) (int, bool) {
	return getParamID(c, "section_id")
}

func getTipID(c *gin.Context) (int, bool) {
	return getParamID(c, "tip_id")
}
