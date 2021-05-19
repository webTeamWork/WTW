package model

const (
	UserTypeAdmin int8 = 1
	UserTypeUser  int8 = 2
)

type User struct {
	UserID     int
	Email      string
	Password   string
	Nickname   string
	Avatar     string
	UserType   int8
	CreateTime int
}
