package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct {

}

func (d DemoController) Demo(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, "admin")
}