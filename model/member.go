package model

// Member 定义会员信息结构体
type Member struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // 主键
	Name     string `gorm:"size:255;column:name" json:"name"`             // 姓名
	Phone    string `gorm:"size:11;column:phone" json:"phone"`            // 手机号
	Password string `gorm:"size:255;column:password" json:"password"`     // 密码
	Email    string `gorm:"size:255;column:email" json:"email"`           // 邮箱
	State    string `gorm:"size:2;column:state" json:"state"`             // 状态
	Info     string `gorm:"size:255;column:info" json:"info"`             // 描述
	Integral int64  `json:"integral;column:integral"`                     // 会员积分
}

// TableName 指定 Member 结构体对应的数据库表名
func (Member) TableName() string {
	return "t_member"
}
