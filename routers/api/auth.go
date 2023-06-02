package api

import (
	"GoGinExample/models"
	"GoGinExample/pkg/e"
	"GoGinExample/pkg/logging"
	"GoGinExample/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// 验证用户名密码并生成token返回
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//首先检测输入是否合法
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	//查找用户是否存在
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	//验证用户用户名密码正确
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			logging.Info(err.Key, err.Message)
			// fmt.Println("<<<<<<<<<<<<")
		}
	}
	//拿到token返回
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
