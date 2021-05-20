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

func ThumbTopic(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)
	err := service.ThumbTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "点赞成功")
}

func FavorTopic(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)
	err := service.FavorTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "收藏成功")
}

func CancelThumbTopic(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)
	err := service.CancelThumbTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "取消点赞成功")
}

func CancelFavorTopic(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)
	err := service.CancelFavorTopic(userID, topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "取消收藏成功")
}
