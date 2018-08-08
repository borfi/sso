package engine

var myEngine Engine

// Engine .
type Engine struct {
	codes map[int]CodeItem
}

func init() {
	myEngine = Engine{
		codes: make(map[int]CodeItem),
	}
}

// New .
func New() *Engine {
	return &myEngine
}
