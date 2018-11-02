package engine

const (
	// Success 成功返回码
	Success = 1

	// SystemException 系统异常
	SystemException = 100
)

func init() {
	//注册系统级返回码
	RegisterCode([]XCode{
		{Code: Success, Msg: "success"},
		{Code: SystemException, Msg: "system exception"},
	})
}
