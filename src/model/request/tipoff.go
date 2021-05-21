package request

type Tipoff struct {
	Content string `json:"content" binding:"required"`
}
