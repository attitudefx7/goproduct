package errno

// 错误码
const (
	ErrParam = 100000 + iota
)

// 错误信息

var messageMap = map[int64]string{
	ErrParam : "参数错误",
}