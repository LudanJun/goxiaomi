package middlewares

import (
	"go_gin_gorm/models"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//中间件 匹配登录权限判断
func InitAdminAuthMiddleware(c *gin.Context) {	
	fmt.Println("InitAdminAuthMiddleware")
	//进行权限判断 没有登录的用户 不能进入后台管理中心

	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653 t后面跟的是随机数 是从js里传过来的
	// c.Request.URL.String() 可以获取访问的地址
	//"/admin/captcha?t=0.37443287758076393" 通过?分割 获取到pathname
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	fmt.Println("pathname:", pathname)//pathname: /admin/captcha
	// 2、获取gin里的Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//判断userinfo里面的信息是否存在
		var userinfoStruct []models.Manager//定义一个切片
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)//将userinfoStr转换成userinfoStruct

		//不为空表示用户已经登录成功
		//取反表示没有登录成功
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			//执行跳转
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")//跳转到登录页面
			}
		}
	} else { //没有登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			//执行跳转
			c.Redirect(302, "/admin/login")
		}
	}

}