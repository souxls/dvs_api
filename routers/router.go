package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Router 注册/api路由
func Router() http.Handler {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g := r.Group("/api")

	v1 := g.Group("/v1")
	{
		// 注册/api/v1/users
		users := v1.Group("users")
		{
			users.GET("", service.users.Query)
			users.GET(":id", service.users.Get)
			users.POST("", service.users.Create)
			users.PUT(":id", service.users.Update)
			users.DELETE(":id", service.users.Delete)
			users.PATCH(":id/enable", service.users.Enable)
			users.PATCH(":id/disable", service.users.Disable)
		}
	}

	return r
}
