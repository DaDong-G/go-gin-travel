package main

import (
	_ "xianyun/Databases"
	"xianyun/Router"
)

func main()  {
	router := Router.InitRouters()
	_ = router.Run("127.0.0.1:8888")
}