package dto

type RoleQueryDTO struct {
	Name   string `json:"name"`  // 名字
	State  string `json:"state"` // 状态
	RoleId int64  `json:"rid"`   // 状态
}
