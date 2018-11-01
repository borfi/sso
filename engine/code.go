package engine

import "sync"

// XCodes 返回码总单元
type XCodes struct {
	sync.Map
}

// XCode 返回码单元
type XCode struct {
	Code int    //返回码
	Msg  string //给用户看的信息
	Info string //给开发人员看的信息
}

// CodeRegister 注册返回码单元
func CodeRegister(codes []XCode) {
	engine.codes.Register(codes)
}

// Get 获取返回码单元
func (xc *XCodes) Get(code int) *XCode {
	v, find := engine.codes.Load(code)
	if !find {
		return nil
	}

	val, ok := v.(XCode)
	if !ok {
		return nil
	}
	return &val
}

// Register 注册返回码单元
func (xc *XCodes) Register(codes []XCode) {
	for k := range codes {
		engine.codes.Store(codes[k].Code, codes[k])
	}
}
