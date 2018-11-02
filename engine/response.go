package engine

// 组装json输出返回
func responseJSON(ctx Context, data interface{}, xerr Error) *XResponse {
	// error
	if xerr != nil {
		return &XResponse{
			Code: xerr.Code(),
			Msg:  xerr.Msg(),
			Info: xerr.Info(),
			Data: nil,
		}
	}

	//success
	r := GetCode(Success)
	return &XResponse{
		Code: r.Code,
		Msg:  r.Msg,
		Info: r.Info,
		Data: data,
	}
}
