package handler

import (
	"encoding/json"
	"net/http"
	"taeho/mud/errcode"
)

func writeJson(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, code, msg string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	err := errcode.ErrFront{
		Code: code,
		Msg:  msg,
	}
	json.NewEncoder(w).Encode(err)
}
