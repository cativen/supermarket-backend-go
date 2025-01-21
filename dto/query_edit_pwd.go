package dto

type QueryEditPwdDTO struct {
	UserName string `json:"username"` // 用户名
	OldPwd   string `json:"oldPwd"`   // 老密码
	NewPwd   string `json:"newPwd"`   // 新密码
}
