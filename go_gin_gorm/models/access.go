package models
//这个表可以当做是两个表
//一个是模块表，一个是菜单表
//自关联表，通过module_id来关联
type Access struct {
	Id          int
	ModuleName  string //模块名称
	ActionName  string //操作名称
	Type        int    //节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string //路由跳转地址
	ModuleId    int    //此module_id和当前模型的id关联       module_id= 0 表示模块
	Sort        int
	Description string
	Status      int
	AddTime     int
	//自己和自己关联
	//foreignKey:ModuleId外键  references:Id主键还是自己
	AccessItem []Access `gorm:"foreignKey:ModuleId;references:Id"` // 外键关联

	// Checked     bool     `gorm:"-"` // 忽略本字段
} 

func (Access) TableName() string {
	return "access"
}
