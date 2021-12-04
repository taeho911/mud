package handler

import (
	"fmt"
	"net/http"
	"taeho/mud/model"
	"taeho/mud/utils"
)

const (
	SESSION_KEY_COOKIE string = "mud_ses"
	SESSION_KEY_LEN    int    = 32
)

var (
	sessionMap = make(map[string]model.Session)
)

func getSession(r *http.Request) (model.Session, error) {
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return model.Session{}, err
	}
	session, exist := sessionMap[cookie.Value]
	if !exist {
		return session, fmt.Errorf("no such session = %v", cookie.Value)
	}
	return session, nil
}

func validateSession(r *http.Request) error {
	ip, err := getIP(r)
	if err != nil {
		return err
	}
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return err
	}
	if sessionMap[cookie.Value].IP != ip {
		return fmt.Errorf("invalid ip address")
	}
	return nil
}

func newSession(w http.ResponseWriter, r *http.Request, username string) error {
	ip, err := getIP(r)
	if err != nil {
		return err
	}
	sessionKey, err := makeSessionKey()
	if err != nil {
		return err
	}
	session := model.Session{
		IP:       ip,
		Username: username,
	}
	session.SetMaketime()
	sessionMap[sessionKey] = session
	http.SetCookie(w, &http.Cookie{
		Name:    SESSION_KEY_COOKIE,
		Value:   sessionKey,
		Expires: session.GetExpirationTime(),
	})
	return nil
}

func makeSessionKey() (string, error) {
	var sessionKey string
	for {
		randomBytes, err := utils.MakeRandom(SESSION_KEY_LEN)
		if err != nil {
			return sessionKey, err
		}
		sessionKey = utils.EncodeBase64(randomBytes)
		if !isExistKey(sessionKey) {
			break
		}
	}
	return sessionKey, nil
}

func isExistKey(sessionKey string) bool {
	_, exist := sessionMap[sessionKey]
	return exist
}
