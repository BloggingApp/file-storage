package service

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrTypeIsNotValid = errors.New("provided an invalid type")
	ErrFileIsNotAnImage = errors.New("provided file is not an image")
	ErrFileIsNotAVideo = errors.New("provided file is not a video")
	ErrFileMustHaveAValidExtension = errors.New("file must have a valid extension")
)
