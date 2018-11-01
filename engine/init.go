package engine

const (
	// SystemException 系统异常
	SystemException = 100
)

func init() {
	// 启动引擎
	start()

	//注册
	CodeRegister([]XCode{
		{Code: SystemException, Msg: "系统异常"},
	})
}
