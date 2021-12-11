package errors

// sign.go는 sign handler용 에러코드
// 2000 ~ 2999
const (
	INVALID_USERNAME      string = "2000"
	INVALID_PASSWORD      string = "2001"
	WRONG_USR_OR_PWD      string = "2002"
	CREATE_SESSION_FAILED string = "2003"
	EXISTING_USERNAME     string = "2004"
)
