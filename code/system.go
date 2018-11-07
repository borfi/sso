package code

import "sso/xengine/xcode"

const (
	// Success 成功返回码
	Success = 1

	// SystemException 系统异常
	SystemException = 100

	// AnalysisConfigError 解析配置文件失败
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
