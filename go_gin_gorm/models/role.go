package models

// 角色表
type Role struct {
    Id       	int    `json:"id"`
	Title	 	string `json:"title"`
	Description string `json:"description"`
	Status 		int    `json:"status"`
	AddTime 	int    `json:"add_time"`
}

func (Role) TableName() string{
	return "role"
}