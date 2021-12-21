package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	// "net/http"
	"fmt"
	// "sms/database"
	// "sms/controller"
)

func Service(ServiceInfo, DatabaseInfo map[string]string){
	router := gin.Default()
	router.POST("/Register", func(c *gin.Context) {
		 // 获取原始字节
		 d, err := c.GetRawData()
		 if err != nil{
			log.Printf(fmt.Sprint(err))
		 }
		Name := gjson.Get(string(d),"name").Str
		Passwd := gjson.Get(string(d),"passwd").Str
		NickName := gjson.Get(string(d),"nickName").Str
		fmt.Println(Name,Passwd,NickName)
	})
	router.Run(":" + ServiceInfo["Port"])
}