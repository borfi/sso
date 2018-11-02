package engine

// GetCode 获取返回码单元
func GetCode(code int) *XCode {
	v, find := engine.codes.m.Load(code)
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
		engine.codes.m.Store(codes[k].Code, codes[k])
	}
}
