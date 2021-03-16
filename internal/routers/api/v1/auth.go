package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog-service/global"
	"go-blog-service/internal/service"
	"go-blog-service/pkg/app"
	"go-blog-service/pkg/errcode"
)

type Auth struct {

}

func NewAuth() Auth {
	return Auth{}
}

func (Auth) GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetAuth errs: %v", err)
		response.ToErrorResponse(errcode.ErrorAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken errs: %v", err)
		response.ToErrorResponse(errcode.ErrorAuthGenerateToken)
		return
	}

	response.ToResponse(gin.H{"token": token})
}
