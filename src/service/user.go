package service

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"forum/src/utils/jwt"
	"strconv"
	"time"
)

func Register(req *request.Register) error {
	var test int
	err := model.DB.Get(&test, "SELECT user_id FROM user WHERE email = ?", req.Email)
	if err != nil {
		return fmt.Errorf("邮箱已存在")
	}

	// gravatar头像
	a := md5.New()
	_, _ = a.Write([]byte(req.Email))
	avatar := fmt.Sprintf("https://gravatar.loli.net/avatar/%x?d=identicon", a.Sum(nil))

	tx, _ := model.DB.Beginx()
	fmt.Print(avatar)
	result, _ := tx.Exec("INSERT INTO user(email, password, nickname, avatar, user_type, create_time) VALUES(?, ?, ?, ?, ?, ?)",
		req.Email, req.Password, req.Nickname, avatar, model.UserTypeUser, time.Now().Unix(),
	)
	userID, err := result.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("注册用户失败")
	}

	// 添加meta表数据
	metas := []string{"topic_count", "comment_count", "view_count", "thumb_count", "favor_count"}
	for _, v := range metas {
		_, err = tx.NamedExec("INSERT INTO user_meta(user_id, meta_name, meta_value) VALUES(:user_id, :meta_name, :meta_value)",
			model.UserMeta{
				UserID:    int(userID),
				MetaName:  v,
				MetaValue: "0",
			},
		)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("设置用户属性失败")
		}
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("注册失败")
	}
	return nil
}

func Login(req *request.Login) (string, error) {
	var userID int
	err := model.DB.Get(&userID, "SELECT user_id FROM user WHERE email = ? and password = ?", req.Email, req.Password)
	if err != nil {
		return "", fmt.Errorf("邮箱不存在或密码不正确")
	}

	return jwt.NewJWT(&jwt.CustomClaims{UserID: userID})
}

// 获取用户model
func getUser(userID int) (*model.User, error) {
	detail := model.User{}
	err := model.DB.Get(&detail, "SELECT * FROM user WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	return &detail, nil
}

func GetUserDetail(userID int) (*model.User, error) {
	detail, err := getUser(userID)
	return detail, err
}

func ChangeUserNickname(userID int, newName string) error {
	detail, err := getUser(userID)
	if err != nil {
		return err
	}

	tx, _ := model.DB.Beginx()
	_, _ = tx.Exec("UPDATE user SET nickname = ? WHERE user_id = ?", newName, detail.UserID)
	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("修改用户昵称失败")
	}
	return nil
}

func ChangeUserPassword(userID int, old, new string) error {
	detail, err := getUser(userID)
	if err != nil {
		return err
	}

	if detail.Password != old {
		return fmt.Errorf("原密码错误")
	}

	tx, _ := model.DB.Beginx()
	_, _ = tx.Exec("UPDATE user SET password = ? WHERE user_id = ?", new, detail.UserID)
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("修改用户密码失败")
	}
	return nil
}

func getUserMeta(userID int, name string) (string, error) {
	var value string
	err := model.DB.Get(&value, "SELECT mate_value FROM user_meta WHERE user_id = ? and meta_name = ?", userID, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("用户属性错误")
		} else {
			return "", fmt.Errorf("获取用户属性失败")
		}
	}
	return value, nil
}

func getUserMetaInt(userID int, name string) (int, error) {
	value, err := getUserMeta(userID, name)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("用户属性设置有误")
	}
	return i, nil
}

func GetUserTopicCount(userID int) (int, error) {
	return getUserMetaInt(userID, "topic_count")
}
