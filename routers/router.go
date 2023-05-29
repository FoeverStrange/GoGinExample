package routers

//routers负责配置路由
import (
	"GoGinExample/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//加载全局中间件Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设定Mode（Debug)
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	return r
}
