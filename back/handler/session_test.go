package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestNewSession"
	if err := newSession(w, r, username); err != nil {
		t.Fatalf("err = %v", err)
	}
	key := w.Result().Cookies()[0].Value
	session := sessionMap[key]
	if session.username != username {
		t.Fatalf("session.username = %v", session.username)
	}
}

func TestGetSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestGetSession"
	newSession(w, r, username)

	copyCookieFromResToReq(w, r)
	session, err := getSession(r)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if session.username != username {
		t.Fatalf("session.username = %v", session.username)
	}
}

func TestUpdateSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestUpdateSession"
	newSession(w, r, username)

	copyCookieFromResToReq(w, r)
	beforeSession, _ := getSession(r)
	beforeTime := beforeSession.maketime

	time.Sleep(1 * time.Second)
	if err := updateSession(r); err != nil {
		t.Fatalf("err = %v", err)
	}

	afterSession, _ := getSession(r)
	afterTime := afterSession.maketime
	if !beforeTime.Before(afterTime) {
		t.Fatalf("beforeTime = %v, afterTime = %v", beforeTime, afterTime)
	}
}

func TestDeleteSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestDeleteSession"
	newSession(w, r, username)

	copyCookieFromResToReq(w, r)
	deleteSession(w, r)

	if _, err := getSession(r); err == nil {
		t.Fatalf("deleteSession failed.")
	}
}

func copyCookieFromResToReq(w *httptest.ResponseRecorder, r *http.Request) {
	var cookies []string
	for _, v := range w.Result().Cookies() {
		cookies = append(cookies, v.Name+"="+v.Value)
	}
	r.Header["Cookie"] = cookies
}
