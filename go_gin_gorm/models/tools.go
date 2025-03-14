package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

//时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()//返回当前时间的时间戳
}

//获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
//md5加密
func Md5(str  string) string {
	h:= md5.New()//创建md5对象
	io.WriteString(h,str)//写入要计算的内容
	return  fmt.Sprintf("%x",h.Sum(nil))//返回md5字符串
}

//表示把string转int
func StringToInt(str  string) (int,error) {
 	n,err:=	strconv.Atoi(str) // 将字符串转换为int
	return n ,err
}

//表示把int转string
func IntToString(n int) string {
	str :=strconv.Itoa(n)  //   将int转换为字符串
   return str
}