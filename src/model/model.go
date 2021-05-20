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

const (
	TopicStatusNormal int8 = 1
	TopicStatusBan    int8 = 2
)

type Topic struct {
	TopicID      int    `db:"topic_id"`
	UserID       int    `db:"user_id"`
	Title        string `db:"title"`
	Content      string `db:"content"`
	CreateTime   int    `db:"create_time"`
	CommentTime  int    `db:"comment_time"`
	SectionID    int    `db:"section_id"`
	Status       int8   `db:"status"`
	CommentCount int    `db:"comment_count"`
}

type TopicMeta struct {
	MetaID    int    `db:"meta_id"`
	TopicID   int    `db:"topic_id"`
	MetaName  string `db:"meta_name"`
	MetaValue string `db:"meta_value"`
}

type Section struct {
	SectionID  int    `db:"section_id"`
	Name       string `db:"name"`
	TopicCount int    `db:"topic_count"`
}
