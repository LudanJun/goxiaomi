package admin

import (
	"encoding/json"
	"fmt"
	"go_gin_gorm/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainController struct{} // 定义一个结构体
// 登录成功进入到主页 这个是主页
func (con MainController) Index(c *gin.Context) {
	//获取 userinfo 对应的session
	session := sessions.Default(c)      //获取session
	userinfo := session.Get("userinfo") //  获取userinfo
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//1.获取用户信息
		var userinfoStruct []models.Manager                  //定义一个userinfoMap
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct) //将userinfoStr 转换成 userinfoStruct

		//2.获取所有的权限
		// 定义一个accessList
		accessList := []models.Access{}
		//权限排序,以sort字段 进行DESC倒序排序
		//需要用到自定义预加载SQL 给Preload第二个参数传一个回调函数
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			//这里的db就是models.DB 只不过是从回调函数里获得的
			return db.Order("access.sort DESC") //2.在对AccessItem子查询进行排序
		}).Order("sort DESC").Find(&accessList) //1.顶级模块直接通过Order进行排序,只能对module_id=0的进行排序 查询出所有的权限

		//3.获取当前角色拥有的权限,并把权限id放到一个map对象里面
		roleAccess := []models.RoleAccess{} //定义一个roleAccess
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int) //定义一个map
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId // 把权限id放到map里面
		}
		//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
		for i := 0; i < len(accessList); i++ {
			// 判断当前权限的id是否在角色权限的Map对象中
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				// 如果存在，给当前权限添加checked属性
				accessList[i].Checked = true
			}
			// 判断当前权限是否有子权限
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				// 判断当前权限的id是否在角色权限的Map对象中
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					// 如果存在，给当前权限添加checked属性
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		fmt.Printf("%#v", accessList)
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login") // 跳转到登录页面

	}
	// c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改状态的方法
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table") // 要修改的表名
	field := c.Query("field") // 要修改的字段名

	// status = ABS(0-1)   1

	// status = ABS(1-1)  0

	// 通过Exec执行原生SQL语句
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 公共修改状态的方法
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}
