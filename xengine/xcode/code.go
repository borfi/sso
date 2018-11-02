package xcode

import "sync"

// 定义返回码总单元
type xCodes struct {
	m sync.Map
}

// XCode 返回码单元
type XCode struct {
	Code int    //返回码
	Msg  string //可以给外界看的信息
	Info string //可以给内部看的信息
}

// 返回码总单元
var xcodes = &xCodes{}

// GetCode 获取返回码单元
func GetCode(code int) *XCode {
	v, find := xcodes.m.Load(code)
	if !find {
		return nil
	}

	val, ok := v.(XCode)
	if !ok {
		return nil
	}
	return &val
}

// RegisterCode 注册返回码单元
func RegisterCode(codes []XCode) {
	for k := range codes {
		xcodes.m.Store(codes[k].Code, codes[k])
	}
}
