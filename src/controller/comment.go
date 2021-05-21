package controller

import (
	"forum/src/model/request"
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func CommentTopic(c *gin.Context) {
	d := request.CommentTopic{}
	if bindRequest(c, &d) != nil {
		apiInputErr(c)
		return
	}

	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)

	err := service.CommentTopic(userID, topicID, &d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "回帖成功")
}
