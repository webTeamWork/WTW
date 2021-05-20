package controller

import (
	"forum/src/model/request"
	"forum/src/service"
	"strconv"

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

func ThumbTopic(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("topic_id"))
	if err != nil {
		apiInputErr(c)
		return
	}

	userID := getUserID(c)
	err = service.ThumbTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "点赞成功")
}

func FavorTopic(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("topic_id"))
	if err != nil {
		apiInputErr(c)
		return
	}

	userID := getUserID(c)
	err = service.FavorTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "收藏成功")
}
