package middlewares

import (
	"encoding/json"
	"fmt"
	"go_gin_gorm/models"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 中间件 匹配登录权限判断
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
	fmt.Println("pathname:", pathname) //pathname: /admin/captcha
	// 2、获取gin里的Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//userinfo:[{"Id":1,"Username":"admin","Password":"e10adc3949ba59abbe56e057f20f883e","Mobile":"15201686411","Email":"518864@qq.com","Status":1,"RoleId":1,"AddTime":1741696650,"IsSuper":1,"Role":{"id":0,"title":"","description":"","status":0,"add_time":0}}]
	fmt.Printf("userinfo:%v", userinfo)
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string) // 将userinfo转换成string类型

	if ok {
		//判断userinfo里面的信息是否存在
		var userinfoStruct []models.Manager                         //定义一个切片
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct) //将userinfoStr转换成userinfoStruct

		//不为空表示用户已经登录成功
		//取反表示没有登录成功
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			//执行跳转
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login") //跳转到登录页面
			}
		} else {
			//用户登录成功 权限判断
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			fmt.Printf("urlPath:%v", urlPath)
			// 判断是否是超级管理员 访问地址不在权限列表里
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				roleAccess := []models.RoleAccess{}
				// 根据角色id获取权限列表
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int) //定义一个map
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId // 把权限id放在map里面
				}
				//2.获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				// pathname   /admin/manager
				access := models.Access{}
				models.DB.Where("url = ?", urlPath).Find(&access)
				//3.判断当前访问的url对应的权限id,是否在权限列表的id中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort() //终止请求
				}
			}
		}
	} else { //没有登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			//执行跳转
			c.Redirect(302, "/admin/login")
		}
	}

}

// 排除权限判断的方法
// 判断urlPath是否在excludeAuthPathSlice里面
func excludeAuthPath(urlPath string) bool {
	//加载配置文件
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	fmt.Println(excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
