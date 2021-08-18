package model

import (
	uuid "github.com/satori/go.uuid"

	"admin/global"
)

type User struct {
	EASBase
	UUID      uuid.UUID `json:"uuid" gorm:"column:uuid"`                                          // 用户UUID
	Username  string    `json:"userName" gorm:""`                                                 // 用户登录名
	Password  string    `json:"-" gorm:""`                                                        // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户"`                                     // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png"` // 用户头像
	// Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId"`
	AuthorityId string `json:"authorityId" gorm:"default:888"` // 用户角色ID
	// Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Login() (user User, err error) {
	db := global.EASMySql.Model(u)

	db = db.Where("username = ?", u.Username).Where("password = ?", u.Password)
	err = db.First(&user).Error

	return
}
