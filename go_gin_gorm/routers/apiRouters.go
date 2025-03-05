package routers

import (
	"006_gorm/controllers/api"

	"github.com/gin-gonic/gin"
)

// 初始化API路由
func ApiRoutersInit(r *gin.Engine) {
	// 创建一个路由组，路径为/api
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.Userlist)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
	}

}
