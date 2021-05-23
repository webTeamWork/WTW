package controller

import (
	"forum/src/model"
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

func GetTipoffList(targetType int8) func(*gin.Context) {
	return func(c *gin.Context) {
		d := request.Pager{}
		if err := bindRequest(c, &d); err != nil {
			return
		}

		var data []model.Tipoff
		var err error
		switch targetType {
		case model.TipoffTargetTypeTopic:
			data, err = service.GetTopicTipoffList(d.Page, d.Limit)
		case model.TipoffTargetTypeComment:
			data, err = service.GetCommentTipoffList(d.Page, d.Limit)
		}

		if err != nil {
			apiErr(c, err.Error())
			return
		}

		list := make([]gin.H, len(data))
		for i, v := range data {
			list[i] = gin.H{
				"tip_id":      v.TipID,
				"user_id":     v.UserID,
				"content":     v.Content,
				"tipoff_time": v.TipoffTime,
			}
			switch targetType {
			case model.TipoffTargetTypeTopic:
				list[i]["topic_id"] = v.TargetID
			case model.TipoffTargetTypeComment:
				list[i]["comment_id"] = v.TargetID
			}
		}

		var count int
		switch targetType {
		case model.TipoffTargetTypeTopic:
			count = service.GetTopicTipoffCount()
		case model.TipoffTargetTypeComment:
			count = service.GetCommentTipoffCount()
		}

		apiOK(c, gin.H{
			"count": count,
			"list":  list,
		}, "获取举报列表成功")
	}
}

func ProcessTipoff(c *gin.Context) {
	tipID, ok := getTipID(c)
	if !ok {
		return
	}

	err := service.ProcessTipoff(tipID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}
	apiOK(c, gin.H{}, "处理举报成功")
}
