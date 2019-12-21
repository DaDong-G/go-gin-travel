package User

import (
	"fmt"
	"time"
	db "xianyun/Databases"
)

type Login struct{
	Captcha string
	Nickname string
	Password string
	Username string
} 

type User struct {
	Id int `json:"id"`
	UserName string `json:"username"`
	CreateTime time.Time `json:"created_at"`
	UpdateTime time.Time `json:"updated_at"`
	Nickname string `json:"nickname"`
	Password string 
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}

func Get()  {
}


func Delete()  {
}


func (u *User)Create()(int64, error){
	fmt.Println(u.UserName,u.Password)
	ret, err := db.SqlDB.Exec("insert into td_user VALUES (default ,?,?,?,?,?,?,?)",
		u.UserName,u.Password,u.Avatar,u.Email,u.Nickname,u.UpdateTime,u.CreateTime)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	n, errLast := ret.LastInsertId()
	if errLast != nil{
		print("错误")
		return 0,err
	}
	return n,nil
}


func Set(){
	
}