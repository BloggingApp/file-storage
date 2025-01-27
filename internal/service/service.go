package service

import (
	"mime/multipart"

	"go.uber.org/zap"
)

type Uploader interface {
	Upload(typ string, path string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type Service struct {
	Uploader
}

func New(logger *zap.Logger) *Service {
	return &Service{
		Uploader: newUploaderService(logger),
	}
}
