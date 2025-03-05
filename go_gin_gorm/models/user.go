// 用户模型
package models

//操作哪个数据库表就建立哪个数据库表的模型

///数据库表user的模型
type User struct { //默认表名是'users' 但是我们定义的是user 所以需要再下面指定表名
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int   `json:"age"`
	Email    string `json:"email"`
	AddTime  int `json:"add_time"`
}

// TableName函数用于指定表名
// 表示配置操作数据库的表名称
func (User) TableName() string {
	return "user" //指定表名为user
}