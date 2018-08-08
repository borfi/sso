package engine

// CodeItem .
type CodeItem struct {
	Code int
	Msg  string
}

// RegisterCode .
func RegisterCode(codes []CodeItem) {
	for _, v := range codes {
		myEngine.codes[v.Code] = v
	}
}
