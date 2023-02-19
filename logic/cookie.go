package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

//定义结构体  缓存结构体 私有
type ginCookie struct{}

//写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	DesKey := "123x5678"
	decrypt, err := DesEncrypt(bytes, []byte(DesKey))

	if err != nil {
		fmt.Println("123-----")
	}
	c.SetCookie(key, string(decrypt), 3600, "/", "http://127.0.0.1:8080/", false, true)
}

//获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {

	valueStr, err1 := c.Cookie(key)
	if err1 != nil {
		fmt.Println(err1)
	}

	if err1 == nil && valueStr != "" && valueStr != "[]" {
		DesKey := "123x5678"
		decrypt, err1 := DesDecrypt([]byte(valueStr), []byte(DesKey))
		err2 := json.Unmarshal(decrypt, obj)
		return err2 == nil && err1 == nil
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "/", "http://127.0.0.1:8080/", false, true)
}

//实例化结构体
var Cookie = &ginCookie{}
