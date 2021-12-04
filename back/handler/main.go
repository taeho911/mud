package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"taeho/mud/errcode"
)

func writeJson(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, code, msg string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	err := errcode.ErrFront{
		Code: "MUD-ERR-" + code,
		Msg:  msg,
	}
	json.NewEncoder(w).Encode(err)
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
