package service

import (
	// "github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
	// "log"
	// "net/http"
	"fmt"
	// "sms/database"
	// "sms/controller"
	"sms/authentication"
	"time"
)

func Service(ServiceInfo, DatabaseInfo map[string]string) {
	// router := gin.Default()
	// router.POST("/Register", func(c *gin.Context) {
	// 	 // 获取原始字节
	// 	 d, err := c.GetRawData()
	// 	 if err != nil{
	// 		log.Printf(fmt.Sprint(err))
	// 	 }
	// 	Name := gjson.Get(string(d),"name").Str
	// 	Passwd := gjson.Get(string(d),"passwd").Str
	// 	NickName := gjson.Get(string(d),"nickName").Str
	// 	fmt.Println(Name,Passwd,NickName)
	// })
	// router.Run(":" + ServiceInfo["Port"])

	TokenString, err := authentication.Setting(map[string]string{"ID": "39043D95-DDF5-40A7-B8B2-53229F6C6E41", "JwtkeyString": "123"}, time.Now().Add(30*time.Second))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(TokenString)

	ID, err := authentication.Getting(map[string]string{"TokenString": TokenString, "JwtkeyString": "123"})
	if err != nil {
		fmt.Println(err)
	}
	if ID == "" {
		fmt.Println("验证失败")
	} else {
		fmt.Println(ID)
	}
}