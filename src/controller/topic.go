package controller

import (
	"forum/src/model"
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
	d := request.Pager{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	sectionID, ok := getSectionID(c)
	if !ok {
		return
	}

	data, err := service.GetSectionTopicList(sectionID, d.Page, d.Limit)
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

func GetTopicDetail(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}

	topic, err := service.GetTopic(topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	user, _ := service.GetUserDetail(topic.UserID)
	view, _ := service.GetTopicViewCount(topicID)
	thumb, _ := service.GetTopicThumbCount(topicID)
	favor, _ := service.GetTopicFavorCount(topicID)
	canSee := true
	if topic.Status == model.CommentStatusBan || topic.Status == model.TopicStatusDelete {
		canSee = false
		topic.Content = "由于被屏蔽或被删除，该贴子不可见，"
	}

	apiOK(c, gin.H{
		"user_nickname": user.Nickname,
		"title":         topic.Title,
		"content":       topic.Content,
		"create_time":   topic.CreateTime,
		"comment_time":  topic.CommentTime,
		"view_count":    view,
		"thumb_count":   thumb,
		"favor_count":   favor,
		"can_see":       canSee,
	}, "获取帖子详情成功")
}

func Search(c *gin.Context) {
	d := request.Search{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	data, count, err := service.Search(d.Content, d.Page, d.Limit)
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

	apiOK(c, gin.H{
		"count": count,
		"list":  list,
	}, "获取搜索结果成功")
}

func GetUserTopicList(c *gin.Context) {
	d := request.Pager{}
	if err := bindRequest(c, &d); err != nil {
		return
	}

	userID := getUserID(c)
	data, err := service.GetUserTopicList(userID, d.Page, d.Limit)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	type listItem struct {
		TopicID      int    `json:"topic_id"`
		Title        string `json:"title"`
		Introduction string `json:"introduction"`
		CreateTime   int    `json:"create_time"`
		ViewCount    int    `json:"view_count"`
		ThumbCount   int    `json:"thumb_count"`
		FavorCount   int    `json:"favor_count"`
	}

	list := make([]listItem, len(data))
	for i, v := range data {
		viewCount, _ := service.GetTopicViewCount(v.TopicID)
		thumbCount, _ := service.GetTopicThumbCount(v.TopicID)
		favorCount, _ := service.GetTopicFavorCount(v.TopicID)
		var intro string
		if len(v.Content) > 100 {
			intro = v.Content[:100]
		} else {
			intro = v.Content
		}
		list[i] = listItem{
			TopicID:      v.TopicID,
			Title:        v.Title,
			Introduction: intro,
			CreateTime:   v.CreateTime,
			ViewCount:    viewCount,
			ThumbCount:   thumbCount,
			FavorCount:   favorCount,
		}
	}

	count, _ := service.GetUserTopicCount(userID)

	apiOK(c, gin.H{
		"count": count,
		"list":  list,
	}, "获取用户发布帖子列表成功")
}

func GetUserRecordList(recordType int8) func(*gin.Context) {
	return func(c *gin.Context) {
		if recordType != model.RecordTypeView && recordType != model.RecordTypeThumb && recordType != model.RecordTypeFavor {
			apiErr(c, "未知记录类型")
			return
		}

		d := request.Pager{}
		if err := bindRequest(c, &d); err != nil {
			apiInputErr(c)
			return
		}
		userID := getUserID(c)

		records, err := service.GetUserRocordList(userID, recordType, d.Page, d.Limit)
		if err != nil {
			apiErr(c, err.Error())
			return
		}

		type listItem struct {
			TopicID      int    `json:"topic_id"`
			Title        string `json:"title"`
			Introduction string `json:"introduction"`
			ViewCount    int    `json:"view_count"`
		}

		list := make([]listItem, len(records))
		for i, record := range records {
			topic, _ := service.GetTopic(record.TopicID)
			var intro string
			if len(topic.Content) > 100 {
				intro = topic.Content[:100]
			} else {
				intro = topic.Content
			}
			viewCount, _ := service.GetTopicViewCount(topic.TopicID)
			list[i] = listItem{
				TopicID:      i,
				Title:        topic.Title,
				Introduction: intro,
				ViewCount:    viewCount,
			}
		}

		var count int
		switch recordType {
		case model.RecordTypeView:
			count, _ = service.GetUserViewCount(userID)
		case model.RecordTypeThumb:
			count, _ = service.GetUserThumbCount(userID)
		case model.RecordTypeFavor:
			count, _ = service.GetUserFavorCount(userID)
		}

		apiOK(c, gin.H{
			"count": count,
			"list":  list,
		}, "获取用户记录列表成功")
	}
}

func BanTopic(c *gin.Context) {
	topicID, ok := getTopicID(c)
	if !ok {
		return
	}

	err := service.BanTopic(topicID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "屏蔽帖子成功")
}

func BanComment(c *gin.Context) {
	commentID, ok := getCommentID(c)
	if !ok {
		return
	}

	err := service.BanComment(commentID)
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	apiOK(c, gin.H{}, "屏蔽回帖成功")
}
