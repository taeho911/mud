package errcode

type ErrFront struct {
	Code string
	Msg  string
}

// main.go는 공통 에러코드
// 1001 ~ 1999
const (
	INVALID_METHOD       string = "1001"
	INVALID_REQUEST_BODY string = "1002"
)
