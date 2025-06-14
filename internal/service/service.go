package service

import "go.uber.org/zap"

type Uploader interface {
	Upload(d UploadData) (string, error)
	Move(moves map[string]string) error
	Delete(paths []string) error
}

type Service struct {
	Uploader
}

func New(logger *zap.Logger) *Service {
	return &Service{
		Uploader: newUploaderService(logger),
	}
}
