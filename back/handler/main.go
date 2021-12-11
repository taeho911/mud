package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

func writeJson(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, code, msg string, status int) {
	w.WriteHeader(status)
	// w.Header().Set("Content-Type", "application/json")
	// err := model.Err{
	// 	Code: "MUD-ERR-" + code,
	// 	Msg:  msg,
	// }
	// json.NewEncoder(w).Encode(err)
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	w.Write(makeErrStr(code, msg))
}

func parseReqBody(body io.ReadCloser, object interface{}) error {
	marshalledBody, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return err
	}
	if json.Unmarshal(marshalledBody, object); err != nil {
		return err
	}
	return nil
}

func getIP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	return ip, err
}

func makeErrStr(code, msg string) []byte {
	return []byte(fmt.Sprintf("<%v> %v", code, msg))
}
