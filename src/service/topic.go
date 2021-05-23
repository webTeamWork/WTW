package service

import (
	"database/sql"
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"math"
	"strconv"
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

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("发布帖子失败")
	}
	return nil
}

func GetTopic(topicID int) (*model.Topic, error) {
	topic := model.Topic{}
	err := model.DB.Get(&topic, "SELECT * FROM topic WHERE topic_id = ?", topicID)
	if err != nil {
		return nil, fmt.Errorf("帖子不存在")
	}
	return &topic, nil
}

func GetSectionTopicList(sectionID, pi, ps int) ([]model.Topic, error) {
	var list []model.Topic
	err := model.DB.Select(&list, "SELECT *  FROM topic WHERE section_id = ? ORDER BY topic_id DESC LIMIT ?, ?", sectionID, (pi-1)*ps, ps)
	if err != nil {
		return nil, fmt.Errorf("获取帖子列表失败")
	} else if len(list) == 0 {
		return nil, fmt.Errorf("当前页无帖子")
	}
	return list, nil
}

func GetUserTopicList(userID, pi, ps int) ([]model.Topic, error) {
	var list []model.Topic
	err := model.DB.Select(&list, "SELECT *  FROM topic WHERE user_id = ? ORDER BY topic_id DESC LIMIT ?, ?", userID, (pi-1)*ps, ps)
	if err != nil {
		return nil, fmt.Errorf("获取帖子列表失败")
	} else if len(list) == 0 {
		return nil, fmt.Errorf("当前页无帖子")
	}
	return list, nil
}

func getTopicMeta(topicID int, name string) (string, error) {
	var value string
	err := model.DB.Get(&value, "SELECT mate_value FROM topic_meta WHERE topic_id = ? and meta_name = ?", topicID, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("帖子属性错误")
		} else {
			return "", fmt.Errorf("获取帖子属性失败")
		}
	}
	return value, nil
}

func getTopicMetaInt(topicID int, name string) (int, error) {
	value, err := getTopicMeta(topicID, name)
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("帖子属性设置有误")
	}
	return count, nil
}

func GetTopicViewCount(topicID int) (int, error) {
	return getTopicMetaInt(topicID, "view_count")
}
func GetTopicThumbCount(topicID int) (int, error) {
	return getTopicMetaInt(topicID, "thumb_count")
}
func GetTopicFavorCount(topicID int) (int, error) {
	return getTopicMetaInt(topicID, "favor_count")
}

func record(userID, topicID int, recordType int8) error {
	// 获取之前的记录，点赞、收藏不允许重复执行
	r := model.Record{}
	err := model.DB.Get(&r, "SELECT * FROM record WHERE user_id = ? and topic_id = ? and record_type = ?", userID, topicID, recordType)
	if err == nil {
		if recordType == model.RecordTypeThumb || recordType == model.RecordTypeFavor {
			return fmt.Errorf("已经记录过")
		}
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("数据库查询错误")
	}

	// 记录到数据库，或修改浏览时间
	tx, _ := model.DB.Beginx()
	now := int(time.Now().Unix())
	if recordType == model.RecordTypeView {
		_, err = tx.Exec("UPDATE record SET record_time = ? WHERE record_id = ?", now, r.RecordID)
	} else {
		_, err = tx.NamedExec("INSERT INTO record(user_id, record_type, topic_id, record_time) VALUES(:user_id, :record_type, :topic_id, :record_time)",
			model.Record{
				UserID:     userID,
				RecordType: recordType,
				TopicID:    topicID,
				RecordTime: now,
			},
		)
	}
	if err != nil {
		return fmt.Errorf("记录失败")
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("记录失败")
	}
	return nil
}

func ThumbTopic(userID, topicID int) error {
	return record(userID, topicID, model.RecordTypeThumb)
}

func FavorTopic(userID, topicID int) error {
	return record(userID, topicID, model.RecordTypeFavor)
}

func UnRecord(userID, topicID int, recordType int8) error {
	// 获取之前的记录
	r := model.Record{}
	err := model.DB.Get(&r, "SELECT * FROM record WHERE user_id = ? and topic_id = ? and record_type = ?", userID, topicID, recordType)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("记录不存在")
		} else {
			return fmt.Errorf("数据库查询错误")
		}
	}

	// 删除记录
	tx, _ := model.DB.Beginx()
	_, err = tx.Exec("DELETE FROM record WHERE user_id = ? and topic_id = ? and record_type = ?", userID, topicID, recordType)
	if err != nil {
		return fmt.Errorf("删除记录失败")
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("删除记录失败")
	}
	return nil
}

func CancelThumbTopic(userID, topicID int) error {
	return UnRecord(userID, topicID, model.RecordTypeThumb)
}

func CancelFavorTopic(userID, topicID int) error {
	return UnRecord(userID, topicID, model.RecordTypeFavor)
}

func GetUserTopicRecord(userID, topicID int) (thumb, favor bool) {
	thumb, favor = false, false
	record := &model.Record{}
	err := model.DB.Get(&record, "SELECT * FROM record WHERE user_id = ? and topic_id = ? and record_type = ?", userID, topicID, model.RecordTypeThumb)
	if err == nil {
		thumb = true
	}
	err = model.DB.Get(&record, "SELECT * FROM record WHERE user_id = ? and topic_id = ? and record_type = ?", userID, topicID, model.RecordTypeFavor)
	if err == nil {
		favor = true
	}
	return
}

func Search(content string, pi, ps int) ([]model.Topic, int, error) {
	var count int
	err := model.DB.Get(&count, "SELECT count(*) FROM topic WHERE match(title) against(?)", content)
	if err != nil {
		if count == 0 {
			return nil, 0, fmt.Errorf("搜索结果为空")
		} else {
			return nil, 0, fmt.Errorf("统计搜索结果时出现错误")
		}
	}

	if pi > int(math.Ceil(float64(count)/float64(ps))) {
		return nil, count, fmt.Errorf("当前页无帖子数据")
	}

	var list []model.Topic
	err = model.DB.Select(&list, "SELECT * FROM topic WHERE match(title) against(?) ORDER BY topic_id DESC LIMIT ?, ?", content, (pi-1)*ps, ps)
	if err != nil {
		return nil, count, fmt.Errorf("获取搜索结果时出现错误")
	}
	return list, count, nil
}

func GetUserRocordList(userID int, recordType int8, pi, ps int) ([]model.Record, error) {
	var list []model.Record
	err := model.DB.Select(&list, "SELECT *  FROM record WHERE user_id = ? and record_type = ? ORDER BY record_id DESC LIMIT ?, ?", userID, recordType, (pi-1)*ps, ps)
	if err != nil {
		return nil, fmt.Errorf("获取记录列表失败")
	} else if len(list) == 0 {
		return nil, fmt.Errorf("当前页无记录")
	}
	return list, nil
}
