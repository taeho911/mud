package handler

import (
	"fmt"
	"net/http"
	"sync"
	"taeho/mud/utils"
	"time"
)

const (
	SESSION_KEY_COOKIE string        = "mud_ses"
	SESSION_KEY_LEN    int           = 32
	SESSION_TIMEOUT    time.Duration = 30 * time.Minute
)

// ----------------------------------------------------------
// Session Struct
// ----------------------------------------------------------
type session struct {
	ip       string
	username string
	maketime time.Time
}

func (ses *session) isTimeout() bool {
	due := ses.maketime.Add(SESSION_TIMEOUT)
	return time.Now().After(due)
}

func (ses *session) setMaketime() {
	ses.maketime = time.Now()
}

func (ses *session) getExpiryTime() time.Time {
	return ses.maketime.Add(SESSION_TIMEOUT)
}

// ----------------------------------------------------------
// Session Manager Struct
// ----------------------------------------------------------
type sessionManager struct {
	lock sync.RWMutex
	m    map[string]session
}

func (sm *sessionManager) new(w http.ResponseWriter, r *http.Request, username string) error {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	ip, err := getIP(r)
	if err != nil {
		return err
	}

	key, err := sm.makeKey()
	if err != nil {
		return err
	}

	ses := session{
		ip:       ip,
		username: username,
		maketime: time.Now(),
	}
	sm.m[key] = ses
	http.SetCookie(w, &http.Cookie{
		Name:  SESSION_KEY_COOKIE,
		Value: key,
		Path:  "/api",
	})

	return nil
}

func (sm *sessionManager) get(r *http.Request) (session, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return session{}, err
	}

	ses, exist := sm.m[cookie.Value]
	if !exist {
		return ses, fmt.Errorf("no such session = %v", cookie.Value)
	}

	ip, err := getIP(r)
	if err != nil {
		return ses, err
	}
	if ses.ip != ip {
		return ses, fmt.Errorf("invalid ip address = %v", ip)
	}
	if ses.isTimeout() {
		return ses, fmt.Errorf("session timeout")
	}

	return ses, nil
}

func (sm *sessionManager) refresh(r *http.Request) error {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return err
	}

	ses, exist := sm.m[cookie.Value]
	if !exist {
		return fmt.Errorf("no such session = %v", cookie.Value)
	}

	ses.setMaketime()
	sm.m[cookie.Value] = ses

	return nil
}

func (sm *sessionManager) delete(w http.ResponseWriter, r *http.Request) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	cookie, err := r.Cookie(SESSION_KEY_COOKIE)
	if err != nil {
		return
	}

	delete(sm.m, cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   SESSION_KEY_COOKIE,
		MaxAge: -1,
	})
}

func (sm *sessionManager) GC() {
	now := time.Now()
	sm.lock.Lock()
	for key, ses := range sm.m {
		if now.After(ses.getExpiryTime()) {
			delete(sm.m, key)
		}
		delete(sm.m, key)
	}
	sm.lock.Unlock()
}

func (sm *sessionManager) makeKey() (string, error) {
	var key string
	for {
		randomBytes, err := utils.MakeRandom(SESSION_KEY_LEN)
		if err != nil {
			return key, err
		}
		key = utils.UrlEncodeBase64(randomBytes)
		if !sm.isExistKey(key) {
			break
		}
	}
	return key, nil
}

func (sm *sessionManager) isExistKey(key string) bool {
	_, exist := sm.m[key]
	return exist
}

// ----------------------------------------------------------
// Session Manager Instance
// ----------------------------------------------------------
var (
	// ????????? AP????????? ????????? ???????????? ?????? map
	// kubernetes??? ?????? ???????????? ??????????????? sticky session ????????? ??????
	SessionManager = sessionManager{
		lock: sync.RWMutex{},
		m:    make(map[string]session),
	}
)
