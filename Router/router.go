package Router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "xianyun/Controllers"
	. "xianyun/Middlewares"
)


func InitRouters()  *gin.Engine  {
	router := gin.Default()
	// 定义全局中间间
	router.Use(CorsMiddleware())
	{
		router.Static("/assets", "./assets")
		router.StaticFS("/more_static", http.Dir("assets"))
		router.GET("/", IndexApi)

		userGroup := router.Group("/user")
		{
			userGroup.POST("/captcha",Captcha)
			userGroup.POST("/accounts/register",Register)
			userGroup.OPTIONS("/captcha")
			userGroup.OPTIONS("/accounts/register")
		}

		return router
	}
}