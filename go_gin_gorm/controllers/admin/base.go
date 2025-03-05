package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type BaseController struct {
    
}

// 定义一个success方法，用于返回成功信息
func (con BaseController) Success(c *gin.Context,message string,redirectUrl string) {
 
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
	    "message":message, // 返回成功信息
		"redirectUrl":redirectUrl, // 返回重定向URL
	})
}


// 定义一个error方法，用于返回错误信息
func (con BaseController) Error(c *gin.Context,message string,redirectUrl string) {
 
	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
	    "message":message, // 返回成功信息
		"redirectUrl":redirectUrl, // 返回重定向URL
	})
}