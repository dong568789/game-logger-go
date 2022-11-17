package msgdata

type ReqWriteLog struct {
	PlayerID string `form:"playerID" binding:"required"`
	Date     string `form:"date" binding:"required"`
	Token    string `form:"token" binding:"required"`
}

type ReqModel struct {
	Model string   `json:"model"`
	List  []string `json:"list"`
}
