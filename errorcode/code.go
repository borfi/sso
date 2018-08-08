package errorcode

import "sso/engine"

const (
	Success = 100000

	AnalysisConfigError = 501001

	ParamsError = 502001
)

func init() {
	engine.RegisterCode([]engine.CodeItem{
		{Code: Success, Msg: "成功"},

		{Code: AnalysisConfigError, Msg: "解析配置文件失败"},

		{Code: ParamsError, Msg: "参数有误"},
	})
}
