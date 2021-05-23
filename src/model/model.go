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

type UserMeta struct {
	MetaID    int    `db:"meta_id"`
	UserID    int    `db:"user_id"`
	MetaName  string `db:"meta_name"`
	MetaValue string `db:"meta_value"`
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

const (
	RecordTypeView  int8 = 1
	RecordTypeThumb int8 = 2
	RecordTypeFavor int8 = 3
)

type Record struct {
	RecordID   int  `db:"record_id"`
	UserID     int  `db:"user_id"`
	RecordType int8 `db:"record_type"`
	TopicID    int  `db:"topic_id"`
	RecordTime int  `db:"record_time"`
}

const (
	CommentStatusNormal int8 = 1
	CommentStatusBan    int8 = 2
)

type Comment struct {
	CommentID   int    `db:"comment_id"`
	TopicID     int    `db:"topic_id"`
	UserID      int    `db:"user_id"`
	Content     string `db:"content"`
	CommentTime int    `db:"comment_time"`
	Status      int8   `db:"status"`
}

const (
	TipoffTargetTypeTopic   int8 = 1
	TipoffTargetTypeComment int8 = 2
)
const (
	TipoffProcessTypeOpen  int8 = 1
	TipoffProcessTypeClose int8 = 2
)

type Tipoff struct {
	TipID       int    `db:"tip_id"`
	UserID      int    `db:"user_id"`
	TargetType  int8   `db:"target_type"`
	TargetID    int    `db:"target_id"`
	Content     string `db:"content"`
	TipoffTime  int    `db:"tipoff_time"`
	ProcessType int8   `db:"process_type"`
}
