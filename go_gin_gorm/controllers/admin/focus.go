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
	id, err1 := models.Int(c.Query("id")) // 获取id
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
	}
	//上传文件
	focusImg, _ := models.UploadImg(c, "focus_img")

	focus := models.Focus{Id: id} // 根据id查找轮播图
	models.DB.Find(&focus)        // 查找轮播图

	focus.Title = title         // 更新轮播图标题
	focus.FocusType = focusType // 更新轮播图类型
	focus.Link = link           // 更新轮播图链接
	focus.Sort = sort           // 更新轮播图排序
	focus.Status = status       // 更新轮播图状态
	if focusImg != "" {
		focus.FocusImg = focusImg // 如果上传了图片，则更新图片
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+models.String(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}
func (con FocusController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
	} else {
		focus := models.Focus{Id: id}
		models.DB.Delete(&focus)
		//根据自己的需要 要不要删除图片
		// os.Remove("static/upload/20210915/1631694117.jpg")
		con.Success(c, "删除数据成功", "/admin/focus")
	}
}
