package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BrandController struct {
	
}

func (bc BrandController) lists(ctx *gin.Context)  {
	// 接收参数
	// 验证参数
	// 请求数据库



	// 返回结果
	ctx.JSON(http.StatusOK, "data")
}