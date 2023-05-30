package routers

//routers负责配置路由
import (
	"GoGinExample/pkg/setting"
	v1 "GoGinExample/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//加载全局中间件Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设定Mode（Debug)
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取tags
		apiv1.GET("/tags", v1.GetTags)
		//新建tag
		apiv1.POST("/tags", v1.AddTag)
		//修改文章tag
		apiv1.PUT("tags/:id", v1.EditTag)
		//删除文章tag
		apiv1.DELETE("tags/:id", v1.DeleteTag)
	}

	return r
}
