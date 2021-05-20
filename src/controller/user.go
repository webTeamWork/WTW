package controller

import (
	"forum/src/model/request"
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	d := &request.Register{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	err := service.Register(d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "注册成功")
}

func Login(c *gin.Context) {
	d := &request.Login{}
	if err := bindRequest(c, &d); err != nil {
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

func GetUserDetail(c *gin.Context) {
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

func ChangeUserNickname(c *gin.Context) {
	d := &request.ChangeUserNickname{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	userID := c.GetInt("UserID")
	err := service.ChangeUserNickname(userID, d.Nickname)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "修改昵称成功")
}

func ChangeUserPassword(c *gin.Context) {
	d := &request.ChangeUserPassword{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	userID := c.GetInt("UserID")
	err := service.ChangeUserPassword(userID, d.Old, d.New)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "修改密码成功")
}
