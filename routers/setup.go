package routers

import (
	"MiShop/logic"
	"github.com/gin-gonic/gin"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func SetUp(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": logic.UnixToTime,
		"Str2Html":   logic.Str2Html,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	AdminRouters(r)
}
