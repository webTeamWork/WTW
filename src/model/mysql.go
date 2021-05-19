package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// 数据库连接池单例
var DB *sqlx.DB

// 打开数据库连接，执行前应正确加载config，若无法打开数据库会报错
func OpenDatabase() {
	d, err := sqlx.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
			"root", "123456", "127.0.0.1", 3306, "forum",
		),
	)
	if err != nil || d.Ping() != nil {
		panic(fmt.Errorf("打开数据库失败"))
	}
	DB = d
}
