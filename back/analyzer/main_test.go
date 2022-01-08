package analyzer

import "testing"

func TestSend(t *testing.T) {
	msg, err := Send("TestSend")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	t.Log(msg)
}
