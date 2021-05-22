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

func GetUserTopicRecord(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}
	userID := getUserID(c)

	thumb, favor := service.GetUserTopicRecord(userID, topicID)

	apiOK(c, gin.H{
		"thumb": thumb,
		"favor": favor,
	}, "获取记录成功")
}

func GetSectionTopicList(c *gin.Context) {
	pager := request.Pager{}
	if err := bindRequest(c, &pager); err != nil {
		return
	}

	sectionID, ok := getSectionID(c)
	if !ok {
		return
	}

	data, err := service.GetSectionTopicList(sectionID, pager.Page, pager.Limit)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	type listItem struct {
		TopicID      int    `json:"topic_id"`
		UserNickname string `json:"user_nickname"`
		Title        string `json:"title"`
		Introduction string `json:"introduction"`
		CreateTime   int    `json:"create_time"`
		CommentTime  int    `json:"comment_time"`
		ViewCount    int    `json:"view_count"`
	}

	list := make([]listItem, len(data))
	for i, v := range data {
		user, _ := service.GetUserDetail(v.UserID)
		viewCount, _ := service.GetTopicViewCount(v.TopicID)
		var intro string
		if len(v.Content) > 100 {
			intro = v.Content[:100]
		} else {
			intro = v.Content
		}
		list[i] = listItem{
			TopicID:      v.TopicID,
			UserNickname: user.Nickname,
			Title:        v.Title,
			Introduction: intro,
			CreateTime:   v.CreateTime,
			CommentTime:  v.CommentTime,
			ViewCount:    viewCount,
		}
	}

	count, _ := service.GetSectionTopicCount(sectionID)

	apiOK(c, gin.H{
		"count": count,
		"list":  list,
	}, "获取板块页帖子列表成功")
}
