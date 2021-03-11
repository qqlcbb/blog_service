package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog-service/global"
	"go-blog-service/internal/service"
	"go-blog-service/pkg/app"
	"go-blog-service/pkg/convert"
	"go-blog-service/pkg/errcode"
	"go-blog-service/pkg/upload"
)

func NewUpload() Upload {
	return Upload{}
}

type Upload struct {}

func (Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParam.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParam)
		return
	}
	svc := service.NewService(c.Request.Context())

	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile errs: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"file_access_url": fileInfo.AccessUrl})
}
