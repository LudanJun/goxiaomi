package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

// 定义一个名为Index的函数，该函数属于ApiController结构体
func (con ApiController) Index(c *gin.Context) {
	// 返回一个状态码为200的字符串
	c.String(200, "我是一个api接口")
}
// 定义一个名为Userlist的函数，该函数属于ApiController结构体
func (con ApiController) Userlist(c *gin.Context) {
	// 返回一个状态码为200的字符串
	c.String(200, "我是一个api接口-Userlist")
}
// 定义一个名为Plist的函数，该函数属于ApiController结构体
func (con ApiController) Plist(c *gin.Context) {
	// 返回一个状态码为200的字符串
	c.String(200, "我是一个api接口-Plist")
}
