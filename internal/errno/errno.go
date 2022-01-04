package errno

// 定义错误码

type Error struct {
	ErrCode int64 `json:"errCode"`
	ErrMsg string `json:"errMsg"`
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func New(errno int64, errs ...error) error {
	if len(errs) > 0 {
		errMsg := ""
		for _, err := range errs{
			errMsg += err.Error() + ";"
		}
		return &Error{
			ErrCode: errno,
			ErrMsg: errMsg,
		}
	}

	if errMsg, ok := messageMap[errno]; ok {
		return &Error{
			errno,
			errMsg,
		}
	}

	return &Error{
		errno,
		"未知错误",
	}
}