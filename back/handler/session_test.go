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
	if err := SessionManager.new(w, r, username); err != nil {
		t.Fatalf("err = %v", err)
	}
	key := w.Result().Cookies()[0].Value
	session := SessionManager.m[key]
	if session.username != username {
		t.Fatalf("session.username = %v", session.username)
	}
}

func TestGetSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestGetSession"
	SessionManager.new(w, r, username)

	copyCookieFromResToReq(w, r)
	session, err := SessionManager.get(r)
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
	SessionManager.new(w, r, username)

	copyCookieFromResToReq(w, r)
	beforeSession, _ := SessionManager.get(r)
	beforeTime := beforeSession.maketime

	time.Sleep(1 * time.Second)
	if err := SessionManager.refresh(r); err != nil {
		t.Fatalf("err = %v", err)
	}

	afterSession, _ := SessionManager.get(r)
	afterTime := afterSession.maketime
	if !beforeTime.Before(afterTime) {
		t.Fatalf("beforeTime = %v, afterTime = %v", beforeTime, afterTime)
	}
}

func TestDeleteSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestDeleteSession"
	SessionManager.new(w, r, username)

	copyCookieFromResToReq(w, r)
	SessionManager.delete(w, r)

	if _, err := SessionManager.get(r); err == nil {
		t.Fatalf("deleteSession failed.")
	}
}

func TestConcSessionReadAndWrite(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	username := "TestConcSessionReadAndWrite"
	SessionManager.new(w, r, username)

	copyCookieFromResToReq(w, r)

	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				SessionManager.new(httptest.NewRecorder(), r, username)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				SessionManager.get(r)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				SessionManager.refresh(r)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				SessionManager.GC()
			}
		}
	}()

	time.Sleep(5 * time.Second)
	quit <- true
}

func copyCookieFromResToReq(w *httptest.ResponseRecorder, r *http.Request) {
	var cookies []string
	for _, v := range w.Result().Cookies() {
		cookies = append(cookies, v.Name+"="+v.Value)
	}
	r.Header["Cookie"] = cookies
}
