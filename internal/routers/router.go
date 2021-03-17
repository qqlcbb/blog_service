package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog-service/global"
	"go-blog-service/internal/middleware"
	v1 "go-blog-service/internal/routers/api/v1"
	"net/http"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.AccessLog())
	engine.Use(middleware.Recovery())
	engine.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := v1.NewUpload()
	auth := v1.NewAuth()

	engine.POST("/upload/file", upload.UploadFile)
	engine.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 获取token
	engine.GET("/auth", auth.GetAuth)

	group := engine.Group("/api/v1")
	group.Use(middleware.JWT())
	{
		group.POST("/tags", tag.Create)
		group.PUT("/tags/:id", tag.Update)
		group.DELETE("/tags/:id", tag.Delete)
		group.PATCH("/tags/:id", tag.Update)
		group.GET("/tags", tag.List)

		group.POST("/articles", article.Create)
		group.PUT("/articles/:id", article.Update)
		group.DELETE("/articles/:id", article.Delete)
		group.PATCH("/articles/:id", article.Update)
		group.GET("/articles", article.List)
		group.GET("/articles/:id", article.Get)
	}

	return engine
}
