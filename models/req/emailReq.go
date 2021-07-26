package req

type EmailReq struct {
	UserID   int   `json:"user_id" binding:'required'`
	Email    string `json:"email" binding:'email'`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType int `json:"operation_type" binding:'required'`
}

