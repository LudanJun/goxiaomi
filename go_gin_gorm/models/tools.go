package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix() //返回当前时间的时间戳
}

// 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5加密
func Md5(str string) string {
	h := md5.New()                       //创建md5对象
	io.WriteString(h, str)               //写入要计算的内容
	return fmt.Sprintf("%x", h.Sum(nil)) //返回md5字符串
}

// 表示把string转int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str) // 将字符串转换为int
	return n, err
}

// 表示把int转string
func String(n int) string {
	str := strconv.Itoa(n) //   将int转换为字符串
	return str
}

// 获取当前可执行文件所在目录
func execDir() string {
	ex, _ := os.Executable() // 获取执行文件路径
	return filepath.Dir(ex)  // 返回执行文件所在目录
}

// 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	//1.获取上传的文件
	file, err := c.FormFile(picName)
	// fmt.Printf("file---%#v", file)

	if err != nil {
		return "", err
	}

	//2.获取后缀名,判断类型是否正确 .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename) // 获取后缀名
	allowExtMap := map[string]bool{    //定义map类型对象 定义允许上传的文件类型
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("上传文件格式不正确")
	}

	//3.创建图片保存目录 static/upload/20251209
	day := GetDay()
	dir := "./static/upload/" + day
	// 可执行文件同级目录下 static/upload/20251209
	// dir := filepath.Join(execDir(), "static", "upload", day)//这个是存到了项目根目录里再创建了一个tmp文件夹里
	fmt.Println("dir00", dir)
	err1 := os.MkdirAll(dir, 0666) // 创建目录 0755在Mac上这样写权限 0666在Window上这样写  实测这俩都可以
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnix(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName) // 拼接文件保存的目录
	c.SaveUploadedFile(file, dst)   // 将上传的文件保存到指定的目录中
	return dst, nil                 // 返回文件保存的路径
}
