package code

import "sso/engine/xcode"

const (
	// Success .
	Success = 1

	// SystemException .
	SystemException = 100

	// AnalysisConfigError .
	AnalysisConfigError = 101
)

func init() {
	//注册系统级返回码
	xcode.Register([]xcode.Code{
		{Code: Success, Msg: "成功"},
		{Code: SystemException, Msg: "系统异常"},
		{Code: AnalysisConfigError, Msg: "解析配置文件失败"},
	})
}
