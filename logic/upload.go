package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomarkdown/markdown"
	"html/template"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/ini.v1"
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

//获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := strconv.Atoi(config.Section("oss").Key("status").String())
	return ossStatus
}

func UpLoadImg(c *gin.Context, imgName string) (string, error) {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return CosUpLoadImg(c, imgName)
	} else {
		return LocalUpLoadImg(c, imgName)
	}
}

func CosUpLoadImg(c *gin.Context, imgName string) (string, error) {
	formFile, err := c.FormFile(imgName)
	file := formFile
	if err != nil {
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
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName
	dst := path.Join(dir, fileName)
	dst, err = CosUpLoad(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}
func CosUpLoad(file *multipart.FileHeader, dst string) (string, error) {
	//将<bucket>和<region>修改为真实的信息
	//bucket的命名规则为{name}-{appid} ，此处填写的存储桶名称必须为此格式
	u, _ := url.Parse("")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  os.Getenv(""),
			SecretKey: os.Getenv(""),
		},
	})
	f, _ := file.Open()
	_, err := c.Object.Put(context.Background(), dst, f, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return dst, nil
}

func LocalUpLoadImg(c *gin.Context, imgName string) (string, error) {
	formFile, err := c.FormFile(imgName)
	file := formFile
	if err != nil {
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

//格式化输出图片
func FormatImg(str string) string {
	ossStatus := GetOssStatus()
	config, _ := ini.Load("./conf/app.ini")
	if ossStatus == 1 {
		return config.Section("oss").Key("CosDomain").String() + str
	} else {
		return "/" + str
	}
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a float64, num int) float64 {
	return a * float64(num)
}

func SubStr(str string, start int, end int) string {
	len := len(str)
	if start < 0 {
		start = 0
	}
	if start > len {
		start = 0
	}

	if end < 0 {
		end = len
	}
	if end > len {
		end = len
	}
	if start > end {
		start, end = end, start
	}

	return str[start:end]
}

func FormatAttr(str string) string {
	tmpSlice := strings.Split(str, "\n")
	var tmpStr string
	for _, v := range tmpSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		tmpStr += tmpStr + string(output)
	}
	return tmpStr
}
func GetOrderId() string {
	// 2022020312233
	template := "20060102150405"
	return time.Now().Format(template) + GetRandNum(4)
}
