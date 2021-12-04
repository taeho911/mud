package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"taeho/mud/agent"
	"taeho/mud/errcode"
	"taeho/mud/model"
	"taeho/mud/utils"
)

const (
	USER_MAX_USERNAME    int    = 16
	USRE_MIN_PASSWORD    int    = 8
	USRE_MAX_PASSWORD    int    = 16
	USER_USERNAME_REGEXP string = `^\w+$`
	SALT_LEN             int    = 16
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errcode.INVALID_METHOD, "POST only", http.StatusBadRequest)
	}
	var user model.User
	if err := parseReqBody(r.Body, &user); err != nil {
		writeError(w, errcode.INVALID_REQUEST_BODY, "Invalid request body", http.StatusBadRequest)
	}
	if err := validateUsername(user); err != nil {
		writeError(w, errcode.INVALID_USERNAME, err.Error(), http.StatusBadRequest)
	}
	if err := validatePassword(user); err != nil {
		writeError(w, errcode.INVALID_PASSWORD, err.Error(), http.StatusBadRequest)
	}
	salt, err := utils.MakeRandom(SALT_LEN)
	if err != nil {
		writeError(w, errcode.FAILED_POST, err.Error(), http.StatusBadRequest)
	}
	hashedPwd, err := utils.HashPwd([]byte(user.Password), salt)
	if err != nil {
		writeError(w, errcode.FAILED_POST, err.Error(), http.StatusBadRequest)
	}
	if err := agent.SaltInsertOne(r.Context(), &model.Salt{Username: user.Username, Salt: salt}); err != nil {
		writeError(w, errcode.FAILED_POST, err.Error(), http.StatusBadRequest)
	}
	user.Password = utils.EncodeBase64(hashedPwd)
	if err := agent.UserInsertOne(r.Context(), &user); err != nil {
		writeError(w, errcode.FAILED_POST, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, user, http.StatusOK)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errcode.INVALID_METHOD, "POST only", http.StatusBadRequest)
	}
	var user model.User
	if err := parseReqBody(r.Body, &user); err != nil {
		writeError(w, errcode.INVALID_REQUEST_BODY, "Invalid request body", http.StatusBadRequest)
	}
}

func validateUsername(user model.User) error {
	if length := len(user.Username); length == 0 || length > USER_MAX_USERNAME {
		return fmt.Errorf("username's length must be %v ~ %v. input username's length = %v", 1, USER_MAX_USERNAME, length)
	}
	re := regexp.MustCompile(USER_USERNAME_REGEXP)
	if !re.MatchString(user.Username) {
		return fmt.Errorf("username must be combination of [a-zA-Z0-9_]")
	}
	return nil
}

// 패스워드는 아래의 조건을 만족해야 한다.
// 8 <= x <= 16
// 알파벳 대소문자, 숫자, 특수기호의 조합
// 알파뱃 대소문자, 숫자, 특수기호를 각각 하나 이상씩 사용
func validatePassword(user model.User) error {
	if length := len(user.Password); length < USRE_MIN_PASSWORD || length > USRE_MAX_PASSWORD {
		return fmt.Errorf("password's length must be %v ~ %v. input password's length = %v", 1, USER_MAX_USERNAME, length)
	}
	symbolRe := regexp.MustCompile(`[_!@#$%^&*?]`)
	if !symbolRe.MatchString(user.Password) {
		return fmt.Errorf("password should use special symbol at least one")
	}
	var number, lower, upper bool
	for _, v := range user.Password {
		if v >= '0' && v <= '9' {
			number = true
		} else if v >= 'a' && v <= 'z' {
			lower = true
		} else if v >= 'A' && v <= 'Z' {
			upper = true
		}
		if number && lower && upper {
			return nil
		}
	}
	return fmt.Errorf("password should use number, lowercase and uppercase alphabet at least one")
}
