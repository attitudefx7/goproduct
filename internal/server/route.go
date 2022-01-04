package server

import (
	"gitea.programmerfamily.com/go/product/internal/admin"
	"gitea.programmerfamily.com/go/product/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoute(router *gin.Engine)  {
	// 记录开始时间

	// 全局捕获异常
	router.Use(Recover)

	// 记录请求日志
	router.Use(LogRecord)

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "方法不存在")
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "路由不存在")
	})

	registerApi(router)

	registerAdmin(router)

}

// 注册 api接口
func registerApi(router *gin.Engine)  {
	v1 := router.Group("api/v1")

	demoController := new(api.DemoController)

	v1.POST("/demo/name", demoController.Name)
}

// 注册 后台接口
func registerAdmin(router *gin.Engine)  {
	adm := router.Group("/admin")

	demo := new(admin.DemoController)
	adm.GET("/demo/demo", demo.Demo)

	attribute := new(admin.AttributeController)
	adm.POST("/attribute", attribute.Add)
	adm.PUT("/attribute", attribute.Edit)
	adm.DELETE("/attribute", attribute.Delete)
	adm.GET("/attribute", attribute.Lists)
	adm.GET("/attribute/:id", attribute.Lists)

}