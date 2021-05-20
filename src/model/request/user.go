package request

type Register struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangeUserNickname struct {
	Nickname string `json:"nickname" binding:"required"`
}

type ChangeUserPassword struct {
	Old string `json:"old" binding:"required"`
	New string `json:"new" binding:"required"`
}
