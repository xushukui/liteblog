package syserror

//Error 该接口被UnKnowError结构体实现
type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

//New ..
func New(msg string, reason error) UnKnowError {
	err := UnKnowError{}
	err.msg = msg
	err.reason = reason
	return err
}
