package xcode

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// 定义返回码总单元
type xCodes struct {
	m sync.Map
}

// Code 返回码单元
type Code struct {
	Code    int    //返回码
	Msg     string //给外界看的信息
	Info    string //给内部看的信息
	i18nMsg string //国际化信息(给外界看的)
}

// 返回码总单元
var xcodes = &xCodes{}

// Get 获取返回码单元
func Get(code int) *Code {
	v, find := xcodes.m.Load(code)
	if !find {
		return nil
	}

	val, ok := v.(Code)
	if !ok {
		return nil
	}
	return &val
}

// Register 注册返回码单元
func Register(codes []Code) {
	for k := range codes {
		if isExist(codes[k].Code) {
			panic(fmt.Sprintf("Existed code %d", codes[k].Code))
		}
		xcodes.m.Store(codes[k].Code, codes[k])
	}
}

// Translate 翻译返回码信息(国际化)
func Translate(code int, i18nMsg string) error {
	if i18nMsg == "" {
		return errors.Errorf("Msg i18n is empty %d", code)
	}

	c := Get(code)
	if c == nil {
		return errors.Errorf("Not find code %d", code)
	}

	c.i18nMsg = i18nMsg
	xcodes.m.Store(code, c)
	return nil
}

// 判断返回码是否已经存在
func isExist(code int) bool {
	_, find := xcodes.m.Load(code)
	return find
}