package handler

import (
	"net/http"
	"taeho/mud/errcode"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errcode.INVALID_METHOD, "POST only", http.StatusBadRequest)
	}

}

func validateUsername() {

}

func validatePassword() {

}
