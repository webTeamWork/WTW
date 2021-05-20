package service

import (
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"time"
)

func PostTopic(userID int, req *request.PostTopic) error {
	_, err := getSection(req.SectionID)
	if err != nil {
		return err
	}

	now := int(time.Now().Unix())
	tx, _ := model.DB.Beginx()
	result, err := tx.NamedExec(`INSERT INTO topic(user_id, title, content, create_time, comment_time, section_id, status, comment_count)
		VALUES(:user_id, :title, :content, :create_time, :comment_time, :section_id, :status, :comment_count)`,
		model.Topic{
			UserID:       userID,
			Title:        req.Title,
			Content:      req.Content,
			CreateTime:   now,
			CommentTime:  now,
			SectionID:    req.SectionID,
			Status:       model.TopicStatusNormal,
			CommentCount: 0,
		},
	)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("添加帖子失败")
	}
	topicID, _ := result.LastInsertId()

	// 添加meta
	metas := []string{"view_count", "thumb_count", "favor_count"}
	for _, v := range metas {
		_, err = tx.NamedExec("INSERT INTO topic_meta(topic_id, meta_name, meta_value) VALUES(:topic_id, :meta_name, :meta_value)",
			model.TopicMeta{
				TopicID:   int(topicID),
				MetaName:  v,
				MetaValue: "0",
			},
		)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("设置帖子属性失败")
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("发布帖子失败")
	}
	return nil
}
