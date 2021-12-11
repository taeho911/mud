package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"taeho/mud/agent"
	"taeho/mud/model"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	ctx := context.TODO()
	if err := agent.CreateClient(ctx); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}
	defer agent.DeleteClient(ctx)
	agent.CreateIndexes()
	returnCode := m.Run()
	os.Exit(returnCode)
}

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

func TestGetIP(t *testing.T) {
	testcases := []string{"127.0.0.1:5000", "0.0.0.0", "0.0.0"}
	for i, v := range testcases {
		request := http.Request{
			RemoteAddr: v,
		}
		switch i {
		case 0:
			if ip, err := getIP(&request); err != nil {
				t.Fatalf("ip = %v,err = %v", ip, err)
			}
		case 1, 2:
			if ip, err := getIP(&request); err == nil {
				t.Fatalf("ip = %v,err = %v", ip, err)
			}
		}
	}
}
