package model

const (
	UserTypeAdmin int8 = 1
	UserTypeUser  int8 = 2
)

type User struct {
	UserID     int    `db:"user_id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Nickname   string `db:"nickname"`
	Avatar     string `db:"avatar"`
	UserType   int8   `db:"user_type"`
	CreateTime int    `db:"create_time"`
}
