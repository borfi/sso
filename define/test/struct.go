package define

// Response 最终输出给客户端的json结构
type Response struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 给外界看的信息
	Info string      `json:"info"` // 给内部看的信息
	Data interface{} `json:"data"` // 数据
}
