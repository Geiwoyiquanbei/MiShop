package logic

import (
	"MiShop/dao/redis"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store base64Captcha.Store = redis.RedisStore{}

//生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	// 配置验证码信息
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	//创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate 生成随机 id、base64 图像字符串
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}
func VerifyCaptcha(lid, value string) bool {
	if store.Verify(lid, value, false) {
		return true
	} else {
		return false
	}
}
