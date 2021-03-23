package utils

type ErrCode int
type ErrContent struct {
	HttpCode       int
	ErrMsg         string
	CauseStableErr bool
}

const (
	NoLogin           ErrCode = 100000
	NoPermGetPprof    ErrCode = 100001

	// Add your special ErrCode
)

var ErrMsgInfo = map[ErrCode]ErrContent{
	NoLogin: {
		HttpCode: 401,
		ErrMsg:   "登录信息已过期，请重新登录",
	},
	// Add relate explain of ErrCode
}
