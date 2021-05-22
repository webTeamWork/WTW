package request

type Pager struct {
	Page  int `json:"page" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}
