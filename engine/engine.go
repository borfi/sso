package engine

var myEngine Engine

// Engine .
type Engine struct {
	codes map[int]CodeItem //return code
}

func init() {
	myEngine = Engine{
		codes: make(map[int]CodeItem),
	}
}
