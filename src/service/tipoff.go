package service

import (
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"time"
)

func tipoff(userID, targetID int, targetType int8, req *request.Tipoff) error {
	var err error
	switch targetType {
	case model.TipoffTargetTypeTopic:
		_, err = GetTopic(targetID)
	case model.TipoffTargetTypeComment:
		_, err = getComment(targetID)
	default:
		return fmt.Errorf("举报类型非法")
	}
	if err != nil {
		return err
	}

	now := int(time.Now().Unix())
	tx, _ := model.DB.Beginx()
	_, err = tx.NamedExec("INSERT INTO tipoff(user_id, target_type, target_id, content, tipoff_time, process_type) VALUES(:user_id, :target_type, :target_id, :content, :tipoff_time, :process_type)",
		model.Tipoff{
			UserID:      userID,
			TargetType:  targetType,
			TargetID:    targetID,
			Content:     req.Content,
			TipoffTime:  now,
			ProcessType: model.TipoffProcessTypeOpen,
		},
	)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("举报失败")
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("发生错误导致举报失败")
	}

	return nil
}

func TipoffTopic(userID, topicID int, req *request.Tipoff) error {
	return tipoff(userID, topicID, model.TipoffTargetTypeTopic, req)
}

func TipoffComment(userID, commentID int, req *request.Tipoff) error {
	return tipoff(userID, commentID, model.TipoffTargetTypeComment, req)
}

func getTipoffList(targetType int8, pi, ps int) ([]model.Tipoff, error) {
	var list []model.Tipoff
	err := model.DB.Select(&list, "SELECT * FROM tipoff WHERE target_type = ? ORDER BY tip_id DESC LIMIT ?, ?", targetType, (pi-1)*ps, ps)
	if err != nil {
		return nil, fmt.Errorf("获取举报列表失败")
	} else if len(list) == 0 {
		return nil, fmt.Errorf("当前页无举报记录")
	}
	return list, nil
}

func GetTopicTipoffList(pi, ps int) ([]model.Tipoff, error) {
	return getTipoffList(model.TipoffTargetTypeTopic, pi, ps)
}

func GetCommentTipoffList(pi, ps int) ([]model.Tipoff, error) {
	return getTipoffList(model.TipoffTargetTypeComment, pi, ps)
}

func getTipoffCount(targetType int8) int {
	var count int
	err := model.DB.Get(&count, "SELECT count(*) FROM tipoff WHERE target_type = ?", targetType)
	if err != nil {
		return 0
	}
	return count
}

func GetTopicTipoffCount() int {
	return getTipoffCount(model.TipoffTargetTypeTopic)
}

func GetCommentTipoffCount() int {
	return getTipoffCount(model.TipoffTargetTypeComment)
}
