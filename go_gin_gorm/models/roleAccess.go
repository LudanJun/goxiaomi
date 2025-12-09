package models

type RoleAccess struct {
	AccessId int `json:"access_id"` // 权限id
	RoleId   int `json:"role_id"`   // 角色id
}

func (RoleAccess) TableName() string {
	return "role_access"
}
