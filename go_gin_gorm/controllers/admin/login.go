package admin

import (
	"go_gin_gorm/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}
 func (con LoginController) Index(c *gin.Context) {
	//验证md5 是否正确
	fmt.Println(models.Md5("123456"))
	c.HTML(http.StatusOK, "admin/login/login.html",gin.H{})

	
}
//登录验证 
 func (con LoginController) DoLogin(c *gin.Context) {

	captchaId := c.PostForm("captchaId") //获取后天管理验证码id

	username := c.PostForm("username") //获取后台管理用户名
	password := c.PostForm("password") //获取后台管理密码
	//1.验证验证码是否正确
	verifyValue := c.PostForm("verifyValue") // 获取后台管理页面的验证码
	fmt.Println(username,password)
	//调用models.VerifyCaptcha方法验证验证码
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag  {
		//2.查询数据库  判断用户以及密码是否存在
		userinfoList := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username = ? and password = ?",username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			//3.执行登录, 保存用户信息 执行跳转 
			session := sessions.Default(c)
			//注意:session.Set没法直接保存 将用户信息结构体转换为json格式
			userinfoSlice,_ :=json.Marshal(userinfoList)
			session.Set("userinfo", string(userinfoSlice))

			session.Save() //保存session
			con.Success(c,"登录成功","/admin")//登录成功跳转到后台管理首页
		}else {
			con.Error(c,"用户名或密码错误","/admin/login") 
		}

	} else {
		con.Error(c,"验证码验证 失败","/admin/login")
	}

 }

//获取验证码接口
func (con LoginController) Captcha(c *gin.Context) {
	id,b64s,err := models.MakeCaptcha()
	if(err != nil){
		c.String(http.StatusOK, "获取验证码失败")
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId" : id,	//验证码id
		"captchaImage": b64s, //base64编码的图片
	})
}
 
//退出登录接口
func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)//获取session 
	session.Delete("userinfo") //删除session
	session.Save() //保存session
	con.Success(c,"退出成功","/admin/login")
 }