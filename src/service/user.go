package service

import (
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"forum/src/utils/jwt"
	"time"
)

func Register(req *request.Register) error {
	var test int
	err := model.DB.Get(&test, "SELECT user_id FROM user WHERE email = ?", req.Email)
	if err != nil {
		return fmt.Errorf("邮箱已存在")
	}

	tx, _ := model.DB.Beginx()
	// TODO 头像
	_, _ = tx.Exec("INSERT INTO user(email, password, nickname, avatar, user_type, create_time) VALUES(?, ?, ?, ?, ?, ?)",
		req.Email, req.Password, req.Nickname, "", model.UserTypeUser, time.Now().Unix(),
	)

	// TODO 添加meta表数据

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

func UserDetail(userID int) (*model.User, error) {
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
		tx.Rollback()
		return fmt.Errorf("修改用户昵称失败")
	}
	return nil
}
