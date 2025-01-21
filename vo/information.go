package vo

type InformationVo struct {
	ID       int64  `json:"id"`       // 唯一标识
	Sex      string `json:"sex"`      // 性别
	Username string `json:"username"` // 用户名
	NickName string `json:"nickName"` // 昵称
	HeadImg  string `json:"headImg"`  // 头像链接
	Info     string `json:"info"`     // 个人信息
	Address  string `json:"address"`  // 地址
	Email    string `json:"email"`    // 电子邮件
	Age      int    `json:"age"`      // 年龄
	DeptID   int64  `json:"deptId"`   // 部门ID
	IDCard   string `json:"idCard"`   // 身份证号码
	DeptName string `json:"deptName"` // 部门名称
	Phone    string `json:"phone"`    // 电话号码
}

// DetailEmpVo 定义员工详细信息结构体
type DetailEmpVo struct {
	ID         int64    `json:"id"`         // 唯一标识
	Sex        string   `json:"sex"`        // 性别
	IsAdmin    bool     `json:"isAdmin"`    // 是否是管理员
	Username   string   `json:"username"`   // 用户名
	NickName   string   `json:"nickName"`   // 昵称
	Password   string   `json:"password"`   // 密码
	HeadImg    string   `json:"headImg"`    // 头像链接
	State      string   `json:"state"`      // 状态
	Info       string   `json:"info"`       // 个人信息
	CreateBy   string   `json:"createby"`   // 创建者
	IDCard     string   `json:"idCard"`     // 身份证号码
	CreateTime string   `json:"createTime"` // 创建时间
	LeaveTime  string   `json:"leaveTime"`  // 离职时间
	Address    string   `json:"address"`    // 地址
	Email      string   `json:"email"`      // 电子邮件
	Age        int      `json:"age"`        // 年龄
	DeptID     int64    `json:"deptId"`     // 部门ID
	DeptName   string   `json:"deptName"`   // 部门名称
	RoleNames  []string `json:"roleNames"`  // 角色名称集合
	Phone      string   `json:"phone"`      // 电话号码
}
