package service

import (
	"errors"
	"go-blog-service/global"
	"go-blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (svc *Service) UploadFile(FileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContainExt(FileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if !upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if !upload.CheckMaxSize(FileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	return &FileInfo{
		Name: fileName,
		AccessUrl: accessUrl,
	}, nil
}
