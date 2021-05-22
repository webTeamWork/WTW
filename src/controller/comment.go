package controller

import (
	"forum/src/model"
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

func GetTopicCommentList(c *gin.Context) {
	pager := request.Pager{}
	if err := bindRequest(c, &pager); err != nil {
		return
	}

	topicID, ok := getTopicID(c)
	if !ok {
		return
	}

	topic, err := service.GetTopic(topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	data, err := service.GetTopicCommentList(topicID, pager.Page, pager.Limit)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	type listItem struct {
		CommentID    int    `json:"comment_id"`
		UserNickname string `json:"user_nickname"`
		Content      string `json:"content"`
		CommentTime  int    `json:"comment_time"`
		Ban          bool   `json:"ban"`
	}

	list := make([]listItem, len(data))
	for i, v := range data {
		user, _ := service.GetUserDetail(v.UserID)
		list[i] = listItem{
			CommentID:    v.CommentID,
			UserNickname: user.Nickname,
			Content:      v.Content,
			CommentTime:  v.CommentTime,
			Ban:          false,
		}
		if v.Status == model.CommentStatusBan {
			list[i].Ban = true
			list[i].Content = "该回帖内容被禁止显示"
		}
	}

	apiOK(c, gin.H{
		"count": topic.CommentCount,
		"list":  list,
	}, "获取板块页帖子列表成功")
}
