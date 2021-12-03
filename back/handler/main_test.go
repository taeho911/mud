package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"taeho/mud/model"
	"testing"
	"time"
)

func TestParseReqBody(t *testing.T) {
	mockUser := model.User{
		Username: "TestParseReqBody",
		Password: "TestParseReqBody",
		Maketime: time.Now(),
	}
	jsonBytes, err := json.Marshal(mockUser)
	if err != nil {
		t.Fatalf("json.Marshal failed. err = %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(jsonBytes))
	if err != nil {
		t.Fatalf("http.NewRequest failed. err = %v", err)
	}

	var resultUser model.User
	if err := parseReqBody(req.Body, &resultUser); err != nil {
		t.Fatalf("parseReqBody failed. err = %v", err)
	}

	if resultUser.Username != mockUser.Username {
		t.Fatalf("resultUser.Username = %v, mockUser.Username = %v", resultUser.Username, mockUser.Username)
	}
	if resultUser.Password != mockUser.Password {
		t.Fatalf("resultUser.Password = %v, mockUser.Password = %v", resultUser.Password, mockUser.Password)
	}
}
