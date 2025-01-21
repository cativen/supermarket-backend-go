package dto

type QueryMemberDTO struct {
	BaseQuery
	Phone string `json:"phone"` // 电话
	State string `json:"state"` // 状态
	Name  string `json:"name"`  // 会员名字
}
