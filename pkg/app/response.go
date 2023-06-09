package app

import (
	"GoGinExample/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// 返回c.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
