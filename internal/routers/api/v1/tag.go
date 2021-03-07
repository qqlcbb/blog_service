package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog-service/global"
	"go-blog-service/internal/service"
	"go-blog-service/pkg/app"
	"go-blog-service/pkg/errcode"
)

type Tag struct {

}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorCountFail)
		return
	}

	list, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetListFail)
		return
	}

	response.ToResponseList(list, totalRows)
	return
}

func (t Tag) Get(c *gin.Context) {

}

func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)

	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errors)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewService(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)

	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errors)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errors.Errors()...))
		return
	}

	svc := service.NewService(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}



