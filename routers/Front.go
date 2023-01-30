package routers

import (
	"MiShop/controller/front"
	"github.com/gin-gonic/gin"
)

func FrontRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", front.Index)
	}
}
