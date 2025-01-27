package handler

import "errors"

var (
	errNoType = errors.New("upload file type is not provided")
	errNoFile = errors.New("there is no file")
)
