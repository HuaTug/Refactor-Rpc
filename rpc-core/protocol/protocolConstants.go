package protocol

// ProtocolConstants 协议常量
var ProtocolConstants = struct {
	HEADER_TOTAL_LEN int
	MAGIC            uint16
	VERSION          byte
	REQ_LEN          int
}{
	HEADER_TOTAL_LEN: 42,
	MAGIC:            0x00,
	VERSION:          0x1,
	REQ_LEN:          32,
}
