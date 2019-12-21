package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	indexModel "xianyun/Models/index"
)

// 用户主页
func IndexApi(c *gin.Context) {
	var p indexModel.Photo
	//ph,_ := p.GetOne()
	//fmt.Println(ph.Url,ph.Id)
	photos := p.GetAll()
	//var Msg  struct {
	//	total int
	//	data interface{}
	//}
	// 结构体内部如果不是大写，其他地方无法使用。
	//Msg.data = photos
	//Msg.total = len(photos)
	//fmt.Println(Msg)

	data := map[string]interface{}{}
	data["total"] = len(photos)
	data["data"] = photos
	c.JSON(http.StatusOK,data)
}


