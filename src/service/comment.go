package service

import (
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"time"
)

func CommentTopic(userID, topicID int, req *request.CommentTopic) error {
	_, err := GetTopic(topicID)
	if err != nil {
		return err
	}

	now := int(time.Now().Unix())
	tx, _ := model.DB.Beginx()
	_, err = tx.NamedExec("INSERT INTO comment(topic_id, user_id, content, comment_time, status) VALUES(:topic_id, :user_id, :content, :comment_time, :status)",
		model.Comment{
			TopicID:     topicID,
			UserID:      userID,
			Content:     req.Content,
			CommentTime: now,
			Status:      model.CommentStatusNormal,
		},
	)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("回帖失败")
	}

	_, err = tx.Exec("UPDATE topic SET comment_time = ?, comment_count = comment_count + 1 WHERE topic_id = ?", now, topicID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("回帖相关操作失败")
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("发布回帖失败")
	}

	return nil
}

func getComment(commentID int) (*model.Comment, error) {
	comment := model.Comment{}
	err := model.DB.Get(&comment, "SELECT * FROM comment WHERE comment_id = ?", commentID)
	if err != nil {
		return nil, fmt.Errorf("回帖不存在")
	}
	return &comment, nil
}

func GetTopicCommentList(topicID, pi, ps int) ([]model.Comment, error) {
	var list []model.Comment
	err := model.DB.Select(&list, "SELECT * FROM comment WHERE topic_id = ? ORDER BY comment_id LIMIT ?, ?", topicID, (pi-1)*ps, ps)
	if err != nil {
		return nil, fmt.Errorf("获取回帖失败")
	} else if len(list) == 0 {
		return nil, fmt.Errorf("当前页没有回帖")
	}
	return list, nil
}
