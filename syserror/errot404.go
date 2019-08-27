package syserror

//Error404 .
type Error404 struct {
	UnKnowError
}

//Code ..
func (c Error404) Code() int {
	return 1002
}

func (c Error404) Error() string {
	return "非法访问"
}
