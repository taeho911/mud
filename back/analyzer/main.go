package analyzer

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	analHost string = os.Getenv("ANAL_HOST")
	analPort string = os.Getenv("ANAL_PORT")
)

func Send(text string) (string, error) {
	if len(analHost) < 1 {
		analHost = "127.0.0.1"
	}
	if len(analPort) < 1 {
		analPort = "19011"
	}
	addr := analHost + ":" + analPort
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	fmt.Println("1")
	fmt.Fprint(conn, text)
	fmt.Println("2")
	msg, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("3")
	if err != nil {
		return msg, err
	}
	return msg, err
}
