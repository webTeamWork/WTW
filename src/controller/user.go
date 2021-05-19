package controller

import (
	"forum/src/model/request"
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	d := &request.Register{}
	err := c.ShouldBindJSON(d)
	if err != nil {
		apiInputErr(c)
		return
	}

	err = service.Register(d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "注册成功")
}

func Login(c *gin.Context) {
	d := &request.Login{}
	err := c.ShouldBindJSON(d)
	if err != nil {
		apiInputErr(c)
		return
	}

	token, err := service.Login(d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{
		"token": token,
	}, "登录成功")
}

func UserDetail(c *gin.Context) {
	userID := c.GetInt("UserID")

	detail, err := service.UserDetail(userID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{
		"email":       detail.Email,
		"nickname":    detail.Nickname,
		"avatar":      detail.Avatar,
		"create_time": detail.CreateTime,
	}, "获取用户详情成功")
}
