package models

import (
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

//创建store 使用默认内存存储
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver //创建driver 
	//创建字符串验证码 实例化结构体
	driverString := base64Captcha.DriverString{
		Height:          40,	// 高度
		Width:           100, 	// 长度
		NoiseCount:      0,		// 噪点数量
		ShowLineOptions: 2 | 4, // 显示线条的数量
		Length:          4,		// 验证码长度
		Source:          "1234567890",//验证码字符
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},//字体文件
	}
	driver = driverString.ConvertFonts() //将字体文件转换为driver
	c := base64Captcha.NewCaptcha(driver, store) //创建验证码实例
	id, b64s, err = c.Generate() //生成验证码
	// id, b64s, err := c.Generate() //生成验证码 如果是这种:=变量方式接收 
	// func MakeCaptcha() (string, string, error) {}//返回方式就不用定义变量名了只写类型就行
	return id, b64s, err //返回验证码
	
}


// 验证验证码
//id:是上面生成的id  verifyValue:是用户输入的验证码
func VerifyCaptcha(id string, verifyValue string) bool {
	fmt.Println(id, verifyValue)
	//验证验证码
	//如果传入true 则验证成功后删除
	if store.Verify(id, verifyValue, true) { //验证成功
		return true
	} else {
		return false
	}    
}