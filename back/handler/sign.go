package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"taeho/mud/errcode"
	"taeho/mud/model"
)

const (
	USER_MAX_USERNAME    int    = 16
	USRE_MIN_PASSWORD    int    = 8
	USRE_MAX_PASSWORD    int    = 16
	USER_USERNAME_REGEXP string = `^\w+$`
	USER_PASSWORD_REGEXP string = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[_!@#$%^&*?])[\w!@#$%^&*?]+$`
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errcode.INVALID_METHOD, "POST only", http.StatusBadRequest)
	}
	var user model.User
	if err := parseReqBody(r.Body, &user); err != nil {
		writeError(w, errcode.INVALID_REQUEST_BODY, "Invalid request body", http.StatusBadRequest)
	}
	validateUsername(user)
	validatePassword(user)
}

func validateUsername(user model.User) error {
	if length := len(user.Username); length == 0 || length > USER_MAX_USERNAME {
		return fmt.Errorf("username's length must be %v ~ %v. input username's length = %v", 1, USER_MAX_USERNAME, length)
	}
	re := regexp.MustCompile(USER_USERNAME_REGEXP)
	if !re.Match([]byte(user.Username)) {
		return fmt.Errorf("username must be combination of [a-zA-Z0-9_]")
	}
	return nil
}

func validatePassword(user model.User) error {
	if length := len(user.Password); length < USRE_MIN_PASSWORD || length > USRE_MAX_PASSWORD {
		return fmt.Errorf("password's length must be %v ~ %v. input password's length = %v", 1, USER_MAX_USERNAME, length)
	}
	re := regexp.MustCompile(USER_PASSWORD_REGEXP)
	if !re.Match([]byte(user.Password)) {
		return fmt.Errorf("password must be combination of numbers, alphabets and special simbols")
	}
	return nil
}
