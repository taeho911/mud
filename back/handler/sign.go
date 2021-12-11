package handler

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"taeho/mud/agent"
	"taeho/mud/errors"
	"taeho/mud/model"
	"taeho/mud/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_MAX_USERNAME    int    = 16
	USRE_MIN_PASSWORD    int    = 8
	USRE_MAX_PASSWORD    int    = 16
	USER_USERNAME_REGEXP string = `^\w+$`
	SALT_LEN             int    = 16
)

// ----------------------------------------------------------
// Handler Functions
// ----------------------------------------------------------
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost {
		writeError(w, errors.INVALID_METHOD, "post method only", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := parseReqBody(r.Body, &user); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateUsername(user); err != nil {
		writeError(w, errors.INVALID_USERNAME, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validatePassword(user); err != nil {
		writeError(w, errors.INVALID_PASSWORD, err.Error(), http.StatusBadRequest)
		return
	}
	if isExistUsername(ctx, user.Username) {
		writeError(w, errors.EXISTING_USERNAME, "username exists", http.StatusBadRequest)
		return
	}
	salt, err := utils.MakeRandom(SALT_LEN)
	if err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPwd, err := utils.HashPwd([]byte(user.Password), salt)
	if err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusBadRequest)
		return
	}
	if err := agent.SaltInsertOne(ctx, &model.Salt{Username: user.Username, Salt: salt}); err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = utils.EncodeBase64(hashedPwd)
	if err := agent.UserInsertOne(ctx, &user); err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost {
		writeError(w, errors.INVALID_METHOD, "post method only", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := parseReqBody(r.Body, &user); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := isCorrectPwd(ctx, user); err != nil {
		writeError(w, errors.WRONG_USR_OR_PWD, err.Error(), http.StatusBadRequest)
		return
	}
	if err := newSession(w, r, user.Username); err != nil {
		writeError(w, errors.CREATE_SESSION_FAILED, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, model.User{Username: user.Username}, http.StatusOK)
}

func SignConfirmHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		deleteSession(w, r)
		writeError(w, errors.UNAUTHORIZED, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := updateSession(r); err != nil {
		deleteSession(w, r)
		writeError(w, errors.UNAUTHORIZED, err.Error(), http.StatusUnauthorized)
		return
	}
	writeJson(w, model.User{Username: session.username}, http.StatusOK)
}

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
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
		return fmt.Errorf("password's length must be %v ~ %v. input password's length = %v", USRE_MIN_PASSWORD, USER_MAX_USERNAME, length)
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

func isCorrectPwd(ctx context.Context, user model.User) error {
	salt, err := agent.SaltFindByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	dbUser, err := agent.UserFindByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	hashedPwd, err := utils.HashPwd([]byte(user.Password), salt.Salt)
	if err != nil {
		return err
	}
	if dbUser.Password != utils.EncodeBase64(hashedPwd) {
		return fmt.Errorf("wrong username or password")
	}
	return nil
}

func isExistUsername(ctx context.Context, username string) bool {
	if _, err := agent.UserFindByUsername(ctx, username); err == mongo.ErrNoDocuments {
		return false
	}
	return true
}
