package controller

import (
	"forum/src/model/request"
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func PostTopic(c *gin.Context) {
	d := request.PostTopic{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	userID := getUserID(c)
	err := service.PostTopic(userID, &d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "发布帖子成功")
}
