package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog-service/internal/middleware"
	v1 "go-blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()

	group := engine.Group("/api/v1")
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
