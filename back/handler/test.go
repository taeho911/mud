package handler

import (
	"fmt"
	"net/http"
)

// ----------------------------------------------------------
// Handler Functions
// ----------------------------------------------------------
func SessionTestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- SessionTestHandler ---")

	var bodyMap map[string]string
	parseReqBody(r.Body, &bodyMap)

	switch bodyMap["cmd"] {
	case "1":
		if err := SessionManager.new(w, r, bodyMap["username"]); err != nil {
			fmt.Println("err =", err)
		}
	case "2":
		if err := SessionManager.refresh(r); err != nil {
			fmt.Println("err =", err)
		}
	case "3":
		ses, err := SessionManager.get(r)
		if err != nil {
			fmt.Println("err =", err)
		} else {
			fmt.Printf("ip=%v, username=%v, maketime=%v", ses.ip, ses.username, ses.maketime)
		}
	case "4":
		SessionManager.delete(w, r)
	case "5":
		for k, v := range SessionManager.m {
			fmt.Printf("key=%v, val=%v\n", k, v.username)
		}
	}
}
