package admin

import (
	"006_gorm/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	//获取管理员以及管理员对应的角色数据
	managerList := []models.Manager{}            // 定义一个管理员切片
	models.DB.Preload("Role").Find(&managerList) // 预加载角色信息
	fmt.Printf("%#v", managerList)               // 可以查看切片已经所有类型和数据
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

}

// 进入到增加管理员页面后,把角色列表传递过去渲染出来
func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色 切片
	roleList := []models.Role{}
	models.DB.Find(&roleList) // 查询所有角色
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		//渲染到页面
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	//执行增加 获取角色id
	rolesId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/add")
		return
	}
	//去除字符串空格,获取用户名
	//c.PostForm相当于获取的是表单提交的数据 写的是前端的请求接口
	username := strings.Trim(c.PostForm("username"), "")
	password := strings.Trim(c.PostForm("password"), "")
	email := strings.Trim(c.PostForm("email"), "")
	mobile := strings.Trim(c.PostForm("mobile"), "")

	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或密码的长度不符合要求", "/admin/add")
		return
	}
	//判断管理员是否存在
	//实例化一个管理员对象切片
	managerList := []models.Manager{}
	//获取数据库中管理员数据,判断用户名是否存在
	//把用户名传入到数据库中查询,返回一个管理员对象切片
	models.DB.Where("username = ?", username).Find(&managerList)
	if len(managerList) > 0 { // 如果大于0 说明用户名已经存在 直接返回
		con.Error(c, "此管理员已经存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	// 实例化一个管理员对象
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password), //对密码进行加密
		Email:    email,
		Mobile:   mobile,
		RoleId:   rolesId,
		Status:   1,
		AddTime:  int(models.GetUnix()), //返回当前时间的时间戳
	}
	//将数据插入到数据库中
	err2 := models.DB.Create(&manager).Error
	if err2 != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "增加管理员成功", "/admin/manager")

}
func (con ManagerController) Edit(c *gin.Context) {
	//获取管理员id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}
	//获取当前id对应的管理员信息
	manager := models.Manager{Id: id} //这里直接传入
	models.DB.Find(&manager)          //找到对应id的管理员信息

	//获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,  // 把管理员信息数据传入到模板中
		"roleList": roleList, //把角色信息数据传入到模板中
	})
}

// 编辑确认提交 添加接口后要配置路由POST
func (con ManagerController) DoEdit(c *gin.Context) {
	// 获取角色id
	id, err := models.Int(c.PostForm("id")) // 获取表单提交过来的id
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}

	roleId, err2 := models.Int(c.PostForm("role_id")) // 获取表单提交过来的角色id
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	//获取表单传过来的数据
	username := strings.Trim(c.PostForm("username"), "")
	password := strings.Trim(c.PostForm("password"), "")
	email := strings.Trim(c.PostForm("email"), "")
	mobile := strings.Trim(c.PostForm("mobile"), "")
	if len(mobile) > 11 {
		con.Error(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
		return
	}
	//执行修改
	//可以这样写manager := models.Manager{Id: id,RoleId: roleId, Username: username, Password: password, Email: email, Mobile: mobile }
	manager := models.Manager{Id: id}
	models.DB.Find(&manager) // 先查询到这个管理员信息
	//不能修改用户名
	manager.Username = username
	manager.Password = password //修改密码
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId

	//注意:判断密码是否为空,为空表示不修改密码,不为空表示修改密码
	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			//拼接字符串
			con.Error(c, "密码长度不合法,密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
		}
		manager.Password = models.Md5(password) // 修改加密密码
	}
	err3 := models.DB.Save(&manager).Error // 保存修改后的数据
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/manager")

}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	} else {
		//执行删除
		manager := models.Manager{Id: id}          // 先通过id数据查询到这个管理员信息
		models.DB.Delete(manager, id)              // 删除数据
		con.Success(c, "删除数据成功", "/admin/manager") // 跳转到列表页
	}
}
