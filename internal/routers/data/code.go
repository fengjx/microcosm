package data

var (
	RespCodeOk  = RespCode{1, "ok"}
	RespCodeErr = RespCode{0, "error"}
)

type RespCode struct {
	code int
	msg  string
}
