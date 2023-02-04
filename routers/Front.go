package routers

import (
	"MiShop/controller/itying"
	"github.com/gin-gonic/gin"
)

func FrontRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.Index)
	}
}
