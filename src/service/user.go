package service

import (
	"fmt"
	"forum/src/model"
	"forum/src/model/request"
	"time"
)

func Register(req *request.Register) error {
	var test int
	err := model.DB.Get(&test, "SELECT user_id FROM user WHERE email = ?", req.Email)
	if err != nil {
		return fmt.Errorf("邮箱已存在")
	}

	tx, _ := model.DB.Beginx()
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
