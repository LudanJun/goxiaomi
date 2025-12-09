package admin

import (
	"fmt"
	"go_gin_gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {

	focusList := []models.Focus{} //切片
	models.DB.Find(&focusList)    //查询轮播图列表 返回给focusList
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

// DoAdd 增加轮播图
func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title") // 获取表单内容
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := models.Int(c.PostForm("sort")) //将表单内容转换为int类型
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
	}

	//上传文件
	focusImgSrc, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Printf("focusImgSrc---%#v", focusImgSrc)
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err5 := models.DB.Create(&focus).Error // 将数据插入数据库
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}

}
func (con FocusController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}
func (con FocusController) DoEdit(c *gin.Context) {
	c.String(200, "DoEdit")
}
func (con FocusController) Delete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/delete.html", gin.H{})
}
