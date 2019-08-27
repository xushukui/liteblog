package syserror

//NoLogin 处理未登录状态
type NoLogin struct {
	UnKnowError
}

//Code ..
func (c NoLogin) Code() int {
	return 1001
}

func (c NoLogin) Error() string {
	return "请先登录!"
}
