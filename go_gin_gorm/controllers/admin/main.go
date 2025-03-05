package admin

import (

	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}
//登录成功进入到主页 这个是主页
func (con MainController) Index(c *gin.Context) {
	// //获取 userinfo 对应的session
	// session := sessions.Default(c)//获取session
	// userinfo := session.Get("userinfo") //  获取userinfo
	// //类型断言 来判断 userinfo是不是一个string
	// userinfoStr,ok:= userinfo.(string) 
	// if ok {//如果userinfo是一个string
	//    var userinfoMap []models.Manager //定义一个userinfoMap
	//    json.Unmarshal([]byte(userinfoStr), &userinfoMap) //将userinfoStr 转换成 userinfoMap
	// }else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 	    "username":"session不存在",
	// 	})
	// }
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
