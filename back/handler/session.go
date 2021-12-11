package handler

import (
	"fmt"
	"net/http"
	"taeho/mud/utils"
	"time"
)

type session struct {
	ip       string
	username string
	maketime time.Time
}

func (session *session) IsTimeout() bool {
	due := session.maketime.Add(SESSION_TIMEOUT)
	return time.Now().After(due)
}

func (session *session) SetMaketime() {
	session.maketime = time.Now()
}

func (session *session) GetExpiryTime() time.Time {
	return session.maketime.Add(SESSION_TIMEOUT)
}

const (
	SESSION_KEY_COOKIE string        = "mud_ses"
	SESSION_KEY_LEN    int           = 32
	SESSION_TIMEOUT    time.Duration = 30 * time.Minute
)

var (
	// 세션을 AP메모리 상에서 관리하기 위한 map
	// kubernetes의 분산 디플로이 환경에서는 sticky session 설정이 필요
	sessionMap = make(map[string]session)
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
		if err := updateSession(r); err != nil {
			fmt.Println("err =", err)
		}
	case "3":
		session, err := getSession(r)
		if err != nil {
			fmt.Println("err =", err)
		} else {
			fmt.Println("session.IP =", session.ip)
			fmt.Println("session.Username =", session.username)
			fmt.Println("session.Maketime =", session.maketime)
		}
	case "4":
		deleteSession(w, r)
	}
}

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
func newSession(w http.ResponseWriter, r *http.Request, username string) error {
	ip, err := getIP(r)
	if err != nil {
		return err
	}
	sessionKey, err := makeSessionKey()
	if err != nil {
		return err
	}
	session := session{
		ip:       ip,
		username: username,
	}
	session.SetMaketime()
	sessionMap[sessionKey] = session
	http.SetCookie(w, &http.Cookie{
		Name:  SESSION_KEY_COOKIE,
		Value: sessionKey,
	})
	return nil
}

func getSession(r *http.Request) (session, error) {
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return session{}, err
	}
	session, exist := sessionMap[cookie.Value]
	if !exist {
		return session, fmt.Errorf("no such session = %v", cookie.Value)
	}
	ip, err := getIP(r)
	if err != nil {
		return session, err
	}
	if session.ip != ip {
		return session, fmt.Errorf("invalid ip address = %v", ip)
	}
	if session.IsTimeout() {
		return session, fmt.Errorf("session timeout")
	}
	return session, nil
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return
	}
	delete(sessionMap, cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   SESSION_KEY_COOKIE,
		MaxAge: -1,
	})
}

func updateSession(r *http.Request) error {
	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return err
	}
	session, exist := sessionMap[cookie.Value]
	if !exist {
		return fmt.Errorf("no such session = %v", cookie.Value)
	}
	session.SetMaketime()
	sessionMap[cookie.Value] = session
	return nil
}

func makeSessionKey() (string, error) {
	var sessionKey string
	for {
		randomBytes, err := utils.MakeRandom(SESSION_KEY_LEN)
		if err != nil {
			return sessionKey, err
		}
		sessionKey = utils.UrlEncodeBase64(randomBytes)
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
