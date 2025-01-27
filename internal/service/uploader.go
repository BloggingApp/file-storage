package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	DEFAULT_FILE_PATH_PREFIX = "public/"

	IMAGE_FILE_TYPE = "IMAGE"
	VIDEO_FILE_TYPE = "VIDEO"

	FILE_URL_STRING = "%s/%s"
)

type uploaderService struct {
	logger *zap.Logger
}

func newUploaderService(logger *zap.Logger) Uploader {
	return &uploaderService{
		logger: logger,
	}
}

func (s *uploaderService) saveFile(path string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		return "", ErrFileMustHaveAValidExtension
	}

	imgID := uuid.New()
	var filePath string
	path = strings.TrimSpace(path)
	if path != "" {
		dirPath := filepath.Join(DEFAULT_FILE_PATH_PREFIX, path)
		filePath = filepath.Join(dirPath, imgID.String() + ext)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			s.logger.Sugar().Errorf("failed to create directories: %s", err.Error())
			return "", err
		}
	} else {
		filePath = filepath.Join(DEFAULT_FILE_PATH_PREFIX, imgID.String() + ext)
	}

	createdFile, err := os.Create(filePath)
	if err != nil {
		s.logger.Sugar().Errorf("failed to create file: %s", err.Error())
		return "", err
	}
	defer createdFile.Close()

	if _, err := io.Copy(createdFile, file); err != nil {
		s.logger.Sugar().Errorf("failed to copy src: %s", err.Error())
		return "", err
	}

	filePath = strings.ReplaceAll(filePath, "\\", "/")

	imgURL := fmt.Sprintf(FILE_URL_STRING, viper.GetString("app.origin"), filePath)
	return imgURL, nil
}

func (s *uploaderService) Upload(typ string, path string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		s.logger.Sugar().Errorf("error while uploading a file: %s", err.Error())
		return "", err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		s.logger.Sugar().Errorf("error while uploading a file: %s", err.Error())
		return "", err
	}

	typ = strings.ToUpper(strings.TrimSpace(typ))

	if typ == IMAGE_FILE_TYPE {
		if !filetype.IsImage(buff) {
			return "", ErrFileIsNotAnImage
		}
	} else if typ == VIDEO_FILE_TYPE {
		if !filetype.IsVideo(buff) {
			return "", ErrFileIsNotAVideo
		}
	} else {
		return "", ErrTypeIsNotValid
	}

	return s.saveFile(path, file, fileHeader)
}
