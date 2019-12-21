package Middlewares

import (
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin","http://127.0.0.1")
		c.Header("Access-Control-Allow-Methods","DELETE, GET,POST ,OPTIONS")
		c.Header("Access-Control-Allow-Origin","*")
		c.Header("Access-Control-Allow-Credentials","true")
		//s, _ := c.Cookie("1")
		//fmt.Println(s)
		//c.Header("Access-Control-Max-Age","31536000")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		//c.Header("Content-Type", "application/json;charset=UTF-8")
		//c.Set("haha","hehe") //  设置值
		c.Next() // 先去执行请求的函数
		//status := c.Writer.Status()
		//s,_ :=c.Cookie("123")
		//fmt.Println(s)
		//fmt.Println(status)
		//resp := map[string]string{"hello":"world"}
		//c.JSON(200, resp)
	}
}