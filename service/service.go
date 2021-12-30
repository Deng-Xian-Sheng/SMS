package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	// "net/http"
	"fmt"
	"sms/database"
	// "sms/controller"
	"github.com/google/uuid"
	"sms/authentication"
	"time"
)

func Service(ServiceInfo, DatabaseInfo map[string]string) {
	//准备好Jwt需要的key
	JwtkeyString, err := database.CRUDAuthentication(DatabaseInfo, map[string]string{}, "RJwtkey")
	if err != nil {
		fmt.Println(err)
	}
	if JwtkeyString["RowsAffected"] == "0" {
		NewJwtkey, err := uuid.NewRandom()
		if err != nil {
			log.Printf(fmt.Sprint(err))
		}
		_, err = database.CRUDAuthentication(DatabaseInfo, map[string]string{"NewJwtkey": fmt.Sprint(NewJwtkey)}, "UJwtkey")
		if err != nil {
			log.Printf(fmt.Sprint(err))
		}
		JwtkeyString, err = database.CRUDAuthentication(DatabaseInfo, map[string]string{}, "RJwtkey")
		if err != nil {
			fmt.Println(err)
		}
	}

	router := gin.Default()
	router.POST("/Register", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Printf(fmt.Sprint(err))
		}
		Name := gjson.Get(string(d), "name").Str
		Passwd := gjson.Get(string(d), "passwd").Str
		NickName := gjson.Get(string(d), "nickName").Str
		fmt.Println(Name, Passwd, NickName)
	})
	router.Run(":" + ServiceInfo["Port"])

	//生成Token
	TokenString, err := authentication.Setting(map[string]string{"ID": "39043D95-DDF5-40A7-B8B2-53229F6C6E41", "JwtkeyString": JwtkeyString["Jwtkey"]}, time.Now().Add(24*time.Hour))
	if err != nil {
		log.Printf(fmt.Sprint(err))
	}

	//验证Token
	ID, err := authentication.Getting(map[string]string{"TokenString": TokenString, "JwtkeyString": JwtkeyString["Jwtkey"]})
	if err != nil {
		log.Printf(fmt.Sprint(err))
	}

	if ID == "" {
		fmt.Println("验证失败")
	} else {
		fmt.Println(ID)
	}
}
