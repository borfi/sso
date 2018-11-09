package xresponse

import (
	"sso/engine/xcode"
)

// JSON 组装json输出返回
func JSON(ctx xdefine.Context, data interface{}, xerr xdefine.Error) *xdefine.Response {
	// error
	if xerr != nil {
		return &xdefine.Response{
			Code: xerr.Code(),
			Msg:  xerr.Msg(),
			Info: xerr.Info(),
			Data: nil,
		}
	}

	//success
	r := xcode.GetCode(xcode.Success)
	return &xdefine.Response{
		Code: r.Code,
		Msg:  r.Msg,
		Info: r.Info,
		Data: data,
	}
}
