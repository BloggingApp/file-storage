package dto

import (
	"encoding/json"
	"net/http"
)

type BasicResponse struct {
	Ok      bool   `json:"ok"`
	Details string `json:"details"`
}

func Respond(w http.ResponseWriter, code int, resp BasicResponse) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(code)
	w.Write(body)
}
