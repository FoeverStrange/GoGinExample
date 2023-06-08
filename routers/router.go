package routers

//routers负责配置路由
import (
	"GoGinExample/pkg/setting"
	"GoGinExample/pkg/upload"
	"GoGinExample/routers/api"
	v1 "GoGinExample/routers/api/v1"
	"net/http"

	_ "GoGinExample/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//加载全局中间件Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//设定Mode（Debug)
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	{
		//获取tags
		apiv1.GET("/tags", v1.GetTags)
		//新建tag
		apiv1.POST("/tags", v1.AddTag)
		//修改文章tag
		apiv1.PUT("tags/:id", v1.EditTag)
		//删除文章tag
		apiv1.DELETE("tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
