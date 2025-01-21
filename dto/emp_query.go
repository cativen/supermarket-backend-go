package dto

type EmpQueryDTO struct {
	BaseQuery
	Username string `json:"username"` // 用户名
	NickName string `json:"nickName"` // 昵称
	Age      string `json:"age"`      // 年龄
	Address  string `json:"address"`  // 昵称
	Sex      string `json:"sex"`      // 昵称
	DeptId   int64  `json:"deptId"`   // 状态
}
