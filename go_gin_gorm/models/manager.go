package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	//Manager加载角色 RoleId是外键 角色里的id是主键
	//references:Id表示关联角色表里的id字段 设置为主键
	Role Role `gorm:"foreignKey:RoleId";"references:Id"` // 外键关联配置管理员表关联角色表
}

func (Manager) TableName() string { // 定义表名
	return "manager"
}
