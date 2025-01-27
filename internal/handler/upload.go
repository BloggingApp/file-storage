package handler

import (
	"net/http"
	"strings"

	"github.com/BloggingApp/cdn/internal/dto"
)

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	typ := strings.TrimSpace(r.Header.Get("type"))
	if typ == "" {
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

	url, err := h.services.Uploader.Upload(typ, path, file, fileHeader)
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
