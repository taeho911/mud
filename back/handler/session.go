package handler

import (
	"fmt"
	"net/http"
	"taeho/mud/model"
	"taeho/mud/utils"
	"time"
)

const (
	SESSION_KEY_COOKIE string = "mud_ses"
	SESSION_KEY_LEN    int    = 32
)

var (
	// 세션을 AP메모리 상에서 관리하기 위한 map
	// kubernetes의 분산 디플로이 환경에서는 sticky session 설정이 필요
	sessionMap = make(map[string]model.Session)
)

// ----------------------------------------------------------
// Handler Functions
// ----------------------------------------------------------
func SessionTestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- SessionTestHandler ---")

	var bodyMap map[string]string
	parseReqBody(r.Body, &bodyMap)
	for key, val := range bodyMap {
		fmt.Printf("key = %v, val = %v\n", key, val)
	}

	switch bodyMap["cmd"] {
	case "1":
		if err := newSession(w, r, bodyMap["username"]); err != nil {
			fmt.Println("err =", err)
		}
	case "2":
		if err := validateSession(r); err != nil {
			fmt.Println("err =", err)
		}
	case "3":
		session, err := getSession(r)
		if err != nil {
			fmt.Println("err =", err)
		} else {
			fmt.Println("session.IP =", session.IP)
			fmt.Println("session.Username =", session.Username)
			fmt.Println("session.Maketime =", session.Maketime)
		}
	case "4":
		deleteSession(w, r)
	}
}

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
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
		return fmt.Errorf("invalid ip address = %v", ip)
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
		Expires: session.GetExpiryTime(),
	})
	return nil
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return
	}
	delete(sessionMap, cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:    SESSION_KEY_COOKIE,
		Value:   "",
		Expires: time.Now(),
	})
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
