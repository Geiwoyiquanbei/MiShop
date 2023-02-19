package itying

import (
	"MiShop/dao/mysql"
	"MiShop/logger"
	"MiShop/logic"
	"MiShop/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func GetCaptcha(c *gin.Context) {
	id, s, err := logic.CaptMake()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": s,
	})
}

func LoginController(c *gin.Context) {
	//生成随机数
	prevPage := c.Request.Referer()
	c.HTML(http.StatusOK, "itying/pass/login.html", gin.H{
		"prevPage": prevPage,
	})
}

func RegisterStep1(c *gin.Context) {
	c.HTML(http.StatusOK, "itying/pass/register_step1.html", gin.H{})
}

func RegisterStep2(c *gin.Context) {
	sign := c.Query("sign")
	verifyCode := c.Query("verifyCode")

	//检验验证码是否与session中的相同
	verifySession := sessions.Default(c)
	n := verifySession.Get("verifyCode")
	sessionCode, ok := n.(string)
	if sessionCode != verifyCode || !ok {
		c.Redirect(302, "/pass/registerStep1")
	}
	//检验sign是否与数据库中相同
	user_temp := []models.UserTemp{}
	mysql.DB.Where("sign=?", sign).Find(&user_temp)
	if len(user_temp) > 0 {
		c.HTML(http.StatusOK, "itying/pass/register_step2.html", gin.H{
			"phone":      user_temp[0].Phone,
			"verifyCode": verifyCode,
			"sign":       sign,
		})
	} else {
		c.Redirect(302, "/pass/registerStep1")
	}
}

func RegisterStep3(c *gin.Context) {
	sign := c.Query("sign")
	smsCode := c.Query("smsCode")

	//1、验证短信验证码是否正确
	session := sessions.Default(c)
	sessionSmsCode := session.Get("smsCode")
	sessionSmsCodeStr, ok := sessionSmsCode.(string)

	if !ok || smsCode != sessionSmsCodeStr {
		c.Redirect(302, "/pass/registerStep1")
	}
	//2、获取sign 判断sign是否合法

	userTemp := []models.UserTemp{}
	mysql.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.HTML(http.StatusOK, "itying/pass/register_step3.html", gin.H{
			"smsCode": smsCode,
			"sign":    sign,
		})
	} else {
		c.Redirect(302, "/pass/registerStep1")
	}
}

func SendCode(c *gin.Context) {
	phone := c.Query("phone")
	verifyCode := c.Query("verifyCode")
	captchaId := c.Query("captchaId")
	fmt.Println(captchaId, verifyCode)
	if captchaId == "resend" {
		// 1、注册第二个页面发送验证码的时候需要验证图形验证码
		sessionDefault := sessions.Default(c)
		sessionVerifyCode := sessionDefault.Get("verifyCode")
		sessionVerifyCodeStr, ok := sessionVerifyCode.(string)
		if !ok || verifyCode != sessionVerifyCodeStr {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "非法请求",
			})
			return
		}
	} else {
		// 1、验证图形验证码是否正确 保存图形验证码
		if flag := logic.VerifyCaptcha(captchaId, verifyCode); !flag {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "验证码输入错误，请重试",
			})
			return
		}
		//保存图形验证码
		sessionDefault := sessions.Default(c)
		sessionDefault.Set("verifyCode", verifyCode)
		_ = sessionDefault.Save()
	}
	/*
		2、判断手机格式是否合法
				pattern := `^[\d]{11}$`
				reg := regexp.MustCompile(pattern)
				reg.MatchString(phone)
	*/
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号格式不合法",
		})
		return
	}

	//3、验证手机号是否注册过
	userList := []models.User{}
	mysql.DB.Where("phone=?", phone).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号已经注册，请直接登录",
		})
		return
	}
	//4、判断当前ip地址今天发送短信的次数

	ip := c.ClientIP()
	currentDay := logic.GetDay() //20211211
	var sendCount int64
	mysql.DB.Table("user_temp").Where("ip=? AND add_day=?", ip, currentDay).Count(&sendCount)
	if sendCount > 10 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "此ip今天发送短信的次数已经达到上限，请明天再试",
		})
		return
	}
	//5、验证当前手机号今天发送的次数是否合法
	userTemp := []models.UserTemp{}
	//生成短信验证码
	smsCode := logic.GetRandNum(6)

	sign := logic.Md5(phone + currentDay) //签名：主要用于页面跳转传值
	mysql.DB.Where("phone = ? AND add_day=?", phone, currentDay).Find(&userTemp)
	if len(userTemp) > 0 {
		if userTemp[0].SendCount > 2 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "此手机号今天发送短信的次数已经达到上限，请明天再试",
			})
			return
		} else {
			//1、生成短信验证码  发送验证码  调用前面课程的接口
			fmt.Println("----------自己集成发送短信的接口--------")
			fmt.Println(smsCode)
			//2、服务器保持验证码
			session := sessions.Default(c)
			session.Set("smsCode", smsCode)
			err := session.Save()
			if err != nil {
				return
			}
			//3、更新发送短信的次数
			oneUserTemp := models.UserTemp{}
			mysql.DB.Where("id=?", userTemp[0].Id).Find(&oneUserTemp)
			oneUserTemp.SendCount = oneUserTemp.SendCount + 1
			oneUserTemp.AddTime = int(logic.GetUnix())
			mysql.DB.Save(&oneUserTemp)

			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "发送短信成功",
				"sign":    sign,
			})
			return
		}

	} else {
		//1、生成短信验证码  发送验证码  调用前面课程的接口
		fmt.Println("----------自己集成发送短信的接口--------")
		fmt.Println(smsCode)
		//2、服务器保持验证码
		session := sessions.Default(c)
		session.Set("smsCode", smsCode)
		session.Save()

		//3、记录发送短信的次数

		oneUserTemp := models.UserTemp{
			Ip:        ip,
			Phone:     phone,
			SendCount: 1,
			AddDay:    currentDay,
			AddTime:   int(logic.GetUnix()),
			Sign:      sign,
		}
		mysql.DB.Create(&oneUserTemp)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "发送短信成功",
			"sign":    sign,
		})
		return

	}
	c.JSON(200, gin.H{
		"SendCode": "SendCode",
	})
}

func ValidateSmsCode(c *gin.Context) {

	sign := c.Query("sign")
	smsCode := c.Query("smsCode")
	//1、验证数据是否合法

	userTemp := []models.UserTemp{}
	mysql.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "非法请求",
		})
		return
	}

	//2、验证短信验证码是否正确
	sessionDefault := sessions.Default(c)
	sessionSmsCode := sessionDefault.Get("smsCode")
	sessionSmsCodeStr, ok := sessionSmsCode.(string)
	if !ok || smsCode != sessionSmsCodeStr {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "短信验证码输入错误",
		})
		return
	}
	//3、判断验证码有没有过期   15分

	nowTime := logic.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "短信验证码已过期",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "验证码输入正确",
	})
}

func DoRegister(c *gin.Context) {
	//1、获取表单传过来的数据
	sign := c.PostForm("sign")
	smsCode := c.PostForm("smsCode")
	password := c.PostForm("password")
	rpassword := c.PostForm("rpassword")
	//2、验证smsCode是否合法
	session := sessions.Default(c)
	sessionSmsCode := session.Get("smsCode")
	sessionSmsCodeStr, ok := sessionSmsCode.(string)
	if !ok || smsCode != sessionSmsCodeStr {
		c.Redirect(302, "/")
	}
	//3、验证密码是否合法
	if len(password) < 6 || password != rpassword {
		c.Redirect(302, "/")
	}
	//4、验证签名是否合法
	userTemp := []models.UserTemp{}
	mysql.DB.Where("sign=?", sign).Find(&userTemp)
	fmt.Println(userTemp)
	if len(userTemp) > 0 {
		//4、完成注册
		user := models.User{
			Phone:    userTemp[0].Phone,
			Password: logic.Md5(password), //密码要加密
			LastIp:   userTemp[0].Ip,
			AddTime:  int(logic.GetUnix()),
			Status:   1,
		}
		mysql.DB.Create(&user)

		//5、执行登录
		logic.Cookie.Set(c, "userinfo", user)

		c.Redirect(302, "/")

	} else {
		c.Redirect(302, "/")
	}

}

func DoLogin(c *gin.Context) {
	phone := strings.Trim(c.PostForm("phone"), " ")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	captchaVal := c.PostForm("captchaVal")
	//1、验证图形验证码是否合法
	if flag := logic.VerifyCaptcha(captchaId, captchaVal); !flag {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "图形验证码不正确",
		})
		return
	}

	//2、验证用户名密码是否正确
	password = logic.Md5(strings.Trim(password, " "))
	userList := []models.User{}
	mysql.DB.Where("phone = ? AND password = ?", phone, password).Find(&userList)
	if len(userList) > 0 {
		//执行登录
		logic.Cookie.Set(c, "userinfo", userList[0])
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "用户登录成功",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名或者密码错误",
		})
		return
	}
}

func DoLogOut(c *gin.Context) {
	//删除cookie里面的userinfo执行跳转
	logic.Cookie.Remove(c, "userinfo")
	c.Redirect(302, "/")
}
