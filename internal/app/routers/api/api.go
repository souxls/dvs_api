package api

import (
	"github.com/gin-gonic/gin"
	"github.com/souxls/dvs_api/internal/app/routers/api/ctl"
	"go.uber.org/dig"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(
		cDemo *ctl.Demo,
	) error {

		g := app.Group("/api")

		v1 := g.Group("/v1")
		{
			// 注册/api/v1/demos
			gDemo := v1.Group("demos")
			{
				gDemo.GET("", cDemo.Query)
				gDemo.GET(":id", cDemo.Get)
				gDemo.POST("", cDemo.Create)
				gDemo.PUT(":id", cDemo.Update)
				gDemo.DELETE(":id", cDemo.Delete)
				gDemo.PATCH(":id/enable", cDemo.Enable)
				gDemo.PATCH(":id/disable", cDemo.Disable)
			}
		}

		return nil
	})
}
