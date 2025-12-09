package models

import (
	"fmt"
	// "go_gin_gorm/models"
	// "go_gin_gorm/models"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //全局对象 数据库连接池 对象
var err error   //错误对象

func init() {
	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String() // 获取配置文件中mysql的配置
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	// 连接数据库
	// 获取的数据拼接起来
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	//root:用户名密码  gin 数据库名称
	// dsn := "root:12345678@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //查询时显示字段
		SkipDefaultTransaction: true, //默认禁用事务 性能提升30%
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功", DB)

	//没有这个表会自动创建
	DB.AutoMigrate(Manager{})
	DB.AutoMigrate(Role{})
	DB.AutoMigrate(Access{})
	DB.AutoMigrate(RoleAccess{})
	DB.AutoMigrate(Focus{})
}
