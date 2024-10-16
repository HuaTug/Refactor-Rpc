package protocol

// MsgStatus 请求状态
type MsgStatus byte

const (
	SUCCESS MsgStatus = 0
	FAIL    MsgStatus = 1
)

// IsSuccess 判断是否成功
func IsSuccess(code MsgStatus) bool {
	return code == SUCCESS
}
