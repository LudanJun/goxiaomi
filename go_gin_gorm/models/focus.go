package models

type Focus struct {
	Id    int
	Title string //标题
	//FocusType驼峰形式在数据库里以下划线方式显示focus_type
	FocusType int    //类型 1:pc端上显示 2:移动端上显示 3:小程序
	FocusImg  string //图片
	Link      string //链接
	Sort      int    //排序
	Status    int    //状态
	AddTime   int    //添加时间
}

func (Focus) TableName() string { 
	return "focus" // 表名
}
