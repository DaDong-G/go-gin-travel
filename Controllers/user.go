package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	. "xianyun/Config"
	. "xianyun/Databases"
	"xianyun/Models/Jwt"
	userModel "xianyun/Models/User"
)

// 用户验证码
func Captcha(c *gin.Context)  {
	tel := c.PostForm("tel")
	get := RedisDB.Get(tel)
	if get.Val() != ""{
		c.JSON(http.StatusOK,gin.H{"code":"不得重复访问"})
		return
	}
	rand.Seed(time.Now().UnixNano())
	s := fmt.Sprintf("%d%d%d%d%d%d",rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10),rand.Intn(10))
	err := RedisDB.Set(tel, s, time.Second*60).Err()
	if err != nil{
		panic("error")
	}
	c.JSON(http.StatusOK,gin.H{"code":s})
}

// 用户注册
func Register(c *gin.Context)  {
	login := &userModel.Login{}
	err := c.ShouldBind(&login)
	if err != nil {
		fmt.Println("错误")
	}

	User := &userModel.User{
		UserName:login.Username ,
		Password:login.Password,
		UpdateTime:time.Now(),
		CreateTime:time.Now(),
		Nickname:login.Nickname,
	}

	n,err := User.Create()
	User.Id = int(n)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"msg":"新增用户失败"})
		return
	}
	fmt.Println(err)
	token, err := Jwt.CreateToken([]byte(SecretKey), User.UserName, User.Nickname)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"msg":"加密 token失败"})
	}
	var Msg struct{
		Token string `json:"token"`
		Users interface{} `json:"user"`
	}
	Msg.Token = token
	User.Password = ""
	Msg.Users = User
	c.JSON(http.StatusOK,Msg)
}
