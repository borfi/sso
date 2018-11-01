package code

import "sso/engine"

const (
	// AnalysisConfigError .
	AnalysisConfigError = 501001

	// ParamsError .
	ParamsError = 502001
)

func init() {
	engine.CodeRegister([]engine.XCode{
		{Code: AnalysisConfigError, Msg: "解析配置文件失败"},

		{Code: ParamsError, Msg: "参数有误%v", Info: "params error"},
	})
}
