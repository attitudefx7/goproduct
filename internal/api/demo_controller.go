package api

import (
	"gitea.programmerfamily.com/go/product/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct {
	BaseController
}

func (c *DemoController) Name(ctx *gin.Context)  {

	logger.GetLogger().Debug("Trying to hit GET request")
	ctx.JSON(http.StatusOK, "ok")
}