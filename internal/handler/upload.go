package handler

import (
	"net/http"
	"strings"

	"github.com/BloggingApp/file-storage/internal/dto"
	"github.com/BloggingApp/file-storage/internal/service"
)

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	fileType := strings.TrimSpace(r.FormValue("type"))
	if fileType == "" {
		dto.Respond(w, http.StatusBadRequest, dto.BasicResponse{
			Ok: false,
			Details: errNoType.Error(),
		})
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		dto.Respond(w, http.StatusBadRequest, dto.BasicResponse{
			Ok: false,
			Details: errNoFile.Error(),
		})
		return
	}

	path := strings.TrimSpace(r.FormValue("path"))

	url, err := h.services.Uploader.Upload(service.UploadData{
		FileType: fileType,
		Path: path,
		File: file,
		FileHeader: fileHeader,
	})
	if err != nil {
		dto.Respond(w, http.StatusInternalServerError, dto.BasicResponse{
			Ok: false,
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}
