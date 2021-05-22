package request

type PostTopic struct {
	SectionID int    `json:"section_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type Search struct {
	Content string `json:"content" binding:"required"`
	Pager
}
