package admin

import (
	"gitea.programmerfamily.com/go/product/internal/errno"
	"gitea.programmerfamily.com/go/product/internal/form"
	"gitea.programmerfamily.com/go/product/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AttributeController struct {
	
}

// 新增
func (ac AttributeController) Add(ctx *gin.Context)  {
	// 要不要判断重复
	attributeForm := form.AddAttributeForm{}

	if err := ctx.ShouldBind(&attributeForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errno.New(errno.ErrParam))
		return
	}

	repo := repository.NewAttributeRepository()

	if err := repo.Create(attributeForm);err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusCreated,gin.H{"errCode":0,"errMsg":"success"})
}

// 更新
func (ac AttributeController) Edit(ctx *gin.Context)  {
	// 先查询，确认存在，在更新。要不要判断重复
	attributeForm := form.EditAttributeForm{}

	if err := ctx.ShouldBind(&attributeForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errno.New(errno.ErrParam))
		return
	}

	repo := repository.NewAttributeRepository()

	if err := repo.Update(attributeForm);err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusCreated,gin.H{"errCode":0,"errMsg":"success"})
}

// delete
func (ac AttributeController) Delete(ctx *gin.Context)  {
	// 先查询，确认存在，在删除
	attributeForm := form.DeleteAttributeForm{}

	if err := ctx.ShouldBind(&attributeForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errno.New(errno.ErrParam))
		return
	}

	repo := repository.NewAttributeRepository()

	if err := repo.Delete(attributeForm);err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusCreated,gin.H{"errCode":0,"errMsg":"success"})
}

// detail
func (ac AttributeController) Get(ctx *gin.Context)  {
	attributeForm := form.GetAttributeForm{}

	if err := ctx.ShouldBind(&attributeForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errno.New(errno.ErrParam))
		return
	}

	repo := repository.NewAttributeRepository()

	attitude, err := repo.Get(attributeForm)

	if err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusCreated,gin.H{"errCode":0,"errMsg":"success", "data": attitude})
}

// 列表
func (ac AttributeController) Lists(ctx *gin.Context)  {
	// 接收参数，要有分页
	attributeForm := form.ListAttributeForm{}
	// 验证参数
	// 请求数据库
	if err := ctx.ShouldBindQuery(&attributeForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, errno.New(errno.ErrParam))
		return
	}

	repo := repository.NewAttributeRepository()

	attributes,err := repo.GetList(attributeForm)

	if err != nil {
		_ = ctx.AbortWithError(http.StatusOK, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"errCode":0,"errMsg":"success", "data": attributes})
}