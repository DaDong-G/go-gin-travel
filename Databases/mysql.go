package Databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//因为我们需要在其他地方使用SqlDB这个变量，所以需要大写代表public
var SqlDB *sql.DB

func init() {
	var err error
	//SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")
	SqlDB, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/travel_xian_yun?charset=utf8")
	if err != nil {
		//fmt.Printf("mysql connect error %v", err)
		return
	}
	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}
	// 设置最大连接数
	SqlDB.SetMaxOpenConns(10)
	// 设置最大空闲链接数
	SqlDB.SetMaxIdleConns(10)
	fmt.Println("Mysql-链接成功")
}