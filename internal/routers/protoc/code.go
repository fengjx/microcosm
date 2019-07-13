package protoc

var (
	RespCodeOk  = RespCode{1, "ok"}
	RespCodeErr = RespCode{0, "error"}
	RespCodeNotAuth = RespCode{2, "permision denied"}
)

type RespCode struct {
	code int
	msg  string
}
