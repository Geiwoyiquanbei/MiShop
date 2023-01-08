package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"path"
	"strconv"
	"time"
)

func Float(str string) (float64, error) {
	flo, err := strconv.ParseFloat(str, 64)
	return flo, err
}

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
	return time.Now().Unix()
}
func GetUnixNano() int64 {
	return time.Now().UnixNano()
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
func UpLoadImg(c *gin.Context, imgName string) (string, error) {
	formFile, err := c.FormFile(imgName)
	file := formFile
	if err != nil {
		c.String(200, "上传失败")
		return "", err
	}
	extName := path.Ext(file.Filename)
	extNameMap := make(map[string]bool)
	extNameMap[".jpg"] = true
	extNameMap[".png"] = true
	extNameMap[".gif"] = true
	extNameMap[".jpeg"] = true
	if _, ok := extNameMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}
	day := GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		return "", err
	}
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName
	dst := path.Join(dir, fileName)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}

//把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}
