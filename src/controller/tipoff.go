package controller

import (
	"forum/src/model/request"
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func TipoffTopic(c *gin.Context) {
	d := request.Tipoff{}
	if bindRequest(c, &d) != nil {
		apiInputErr(c)
		return
	}

	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)

	err := service.TipoffTopic(userID, topicID, &d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "举报帖子成功")
}

func TipoffComment(c *gin.Context) {
	d := request.Tipoff{}
	if bindRequest(c, &d) != nil {
		apiInputErr(c)
		return
	}

	commentID, ok := getCommentID(c)
	if !ok {
		return
	}
	userID := getUserID(c)

	err := service.TipoffComment(userID, commentID, &d)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "举报回帖成功")
}
