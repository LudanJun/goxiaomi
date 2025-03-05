package routers

import (
	"006_gorm/controllers/itying"

	"github.com/gin-gonic/gin"
)

// DefaultRoutersInit函数用于初始化默认路由
func DefaultRoutersInit(r *gin.Engine) {
	// 创建一个默认路由组
	defaultRouters := r.Group("/")
	{
		// 添加一个GET请求的路由，路径为"/"，对应的处理函数为itying.DefaultController{}.Index
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		// 添加一个GET请求的路由，路径为"/news"，对应的处理函数为itying.DefaultController{}.News
		defaultRouters.GET("/news", itying.DefaultController{}.News)

	}
}
