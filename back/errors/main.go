package errors

// type Err struct {
// 	Code string
// 	Msg  string
// }

// func (err *Err) Error() string {
// 	return fmt.Sprintf("<%s> %s", err.Code, err.Msg)
// }

// main.go는 공통 에러코드
// 1000 ~ 1999
const (
	INVALID_METHOD       string = "1000"
	INVALID_REQUEST_BODY string = "1001"
	FAILED_POST          string = "1002"
	FAILED_GET           string = "1003"
	FAILED_PUT           string = "1004"
	FAILED_DELETE        string = "1005"
	UNAUTHORIZED         string = "1006"
	INVALID_QUERY        string = "1007"
)
