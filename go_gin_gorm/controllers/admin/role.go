package admin

import (
	"go_gin_gorm/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 角色控制器
type RoleController struct {
	BaseController // 基础控制器
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	fmt.Println(roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
	// c.HTML(http.StatusOK, "admin/role/index.html", gin.H{})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

// 执行添加角色
func (con RoleController) DoAdd(c *gin.Context) {
	//获取表单数据
	title := strings.Trim(c.PostForm("title"), " ")             // 去掉前后空格
	description := strings.Trim(c.PostForm("description"), " ") // 去掉前后空格

	if title == "" {
		con.Error(c, "角色名称不能为空", "/admin/role/add")
		return
	}
	//标题不为空给数据增加数据
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix()) //当前时间戳 int64转int
	err := models.DB.Create(&role).Error //创建数据
	if err != nil {
		con.Error(c, "增加角色失败 请重试", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id")) //获取的id 是string类型
	if err != nil {                              //表示获取的id不是数字类型
		con.Error(c, "传入参数错误", "/admin/role")

	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}

}

// 执行修改
func (con RoleController) DoEdit(c *gin.Context) {
	//获取表单数据
	id, err1 := models.Int(c.PostForm("id")) //获取id
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")             // 去掉前后空格
	description := strings.Trim(c.PostForm("description"), " ") // 去掉前后空格

	if title == "" { //判断标题是否为空
		con.Error(c, "角色名称不能为空", "/admin/role/edit") //不然跳转到编辑页面
		return
	}
	role := models.Role{Id: id} //根据id查询要修改的数据
	models.DB.Find(&role)       //查询要修改的数据
	role.Title = title
	role.Description = description
	err2 := models.DB.Save(&role).Error //保存修改的数据
	if err2 != nil {
		con.Error(c, "修改数据失败请重试", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}

}

// 删除
func (con RoleController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id")) //获取id
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/role")

	} else {
		role := models.Role{Id: id} //根据id查询要修改的数据
		models.DB.Delete(&role)     //删除数据
		con.Success(c, "修改数据成功", "/admin/role")
	}

}
