package todo

type TodoItem struct {
	Instruction string
	Status      string
}

type Output interface {
	Output(items ...TodoItem) error
}

type Write interface {
	Write(items ...TodoItem) error
}

type Read interface {
	Read() []TodoItem
}