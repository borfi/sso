package code

import "sso/engine"

const (
	AnalysisConfigError = 501001

	ParamsError = 502001
)

func init() {
	engine.RegisterCode([]engine.CodeItem{
		{Code: AnalysisConfigError, Msg: "解析配置文件失败"},

		{Code: ParamsError, Msg: "参数有误"},
	})
}
