package model

import (
	"time"
)

type SysUserModel struct {
	ID              int64      `form:"id" gorm:"column:id" json:"id"`                                            //
	Email           string     `form:"email" gorm:"column:email" json:"email" binding:"required"`                // 邮箱
	Phone           string     `form:"phone" gorm:"column:phone" json:"phone" binding:"required"`                // 手机号
	EmailVerifiedAt time.Time  `form:"email_verified_at" gorm:"column:email_verified_at" json:"emailVerifiedAt"` // 邮箱验证时间
	AccountName     string     `form:"account_name" gorm:"column:account_name" json:"accountName"`               // 用户登录名
	NickName        string     `form:"nick_name" gorm:"column:nick_name" json:"nickName"`                        // 昵称
	FullName        string     `form:"full_name" gorm:"column:full_name" json:"fullName"`                        // 实名
	Password        string     `form:"password" gorm:"column:password" json:"password"`                          // 密码
	Avatar          string     `form:"avatar" gorm:"column:avatar" json:"avatar"`                                // 头像
	LastLoginAt     time.Time  `form:"last_login_at" gorm:"column:last_login_at" json:"lastLoginAt"`             // 最后登录日期
	LastToken       string     `form:"last_token" gorm:"column:last_token" json:"lastToken"`                     // 最新登录token
	LastIP          string     `form:"last_ip" gorm:"column:last_ip" json:"lastIP"`                              // 最后登录IP
	RoleID          int        `form:"role_id" gorm:"column:role_id" json:"roleID"`                              // 角色ID
	Status          int        `form:"status" gorm:"column:status" json:"status"`                                // 状态: 1=启用 0=禁用
	DepartmentID    int        `form:"department_id" gorm:"column:department_id" json:"departmentID"`            // 部门ID
	CreatedAt       time.Time  `form:"created_at" gorm:"column:created_at" json:"createdAt"`                     //
	UpdatedAt       time.Time  `form:"updated_at" gorm:"column:updated_at" json:"updatedAt"`                     //
	DeletedAt       *time.Time `form:"deleted_at" gorm:"column:deleted_at;default:null" json:"deletedAt"`        // 删除时间 null未删除
	SideMode        string     `form:"side_mode" gorm:"column:side_mode" json:"sideMode"`                        // 用户侧边主题
	BaseColor       string     `form:"base_color" gorm:"column:base_color" json:"baseColor"`                     // 基础颜色
	ActiveColor     string     `form:"active_color" gorm:"column:active_color" json:"activeColor"`               // 活跃颜色
}

func (model *SysUserModel) TableName() string {
	return "sys_users"
}

var ColumnSysUser = struct {
	ID              string
	Email           string
	Phone           string
	EmailVerifiedAt string
	AccountName     string
	NickName        string
	FullName        string
	Password        string
	Avatar          string
	LastLoginAt     string
	LastToken       string
	LastIP          string
	RoleID          string
	Status          string
	DepartmentID    string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	SideMode        string
	BaseColor       string
	ActiveColor     string
}{
	ID:              "id",
	Email:           "email",
	Phone:           "phone",
	EmailVerifiedAt: "email_verified_at",
	AccountName:     "account_name",
	NickName:        "nick_name",
	FullName:        "full_name",
	Password:        "password",
	Avatar:          "avatar",
	LastLoginAt:     "last_login_at",
	LastToken:       "last_token",
	LastIP:          "last_ip",
	RoleID:          "role_id",
	Status:          "status",
	DepartmentID:    "department_id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	SideMode:        "side_mode",
	BaseColor:       "base_color",
	ActiveColor:     "active_color",
}

type SysUserModelOptions struct {
	apply func(*SysUserModel)
}

func NewSysUserModel(opts ...*SysUserModelOptions) *SysUserModel {
	m := &SysUserModel{}
	for _, opt := range opts {
		opt.apply(m)
	}
	return m
}
