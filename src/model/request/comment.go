package request

type CommentTopic struct {
	Content string `json:"content" binding:"required"`
}
