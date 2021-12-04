package model

import "time"

const (
	SESSION_TIMEOUT time.Duration = 30 * time.Minute
)

type Session struct {
	IP       string
	Username string
	Maketime time.Time
}

func (session *Session) IsTimeout() bool {
	due := session.Maketime.Add(SESSION_TIMEOUT)
	return time.Now().After(due)
}

func (session *Session) SetMaketime() {
	session.Maketime = time.Now()
}

func (session *Session) GetExpiryTime() time.Time {
	return session.Maketime.Add(SESSION_TIMEOUT)
}
