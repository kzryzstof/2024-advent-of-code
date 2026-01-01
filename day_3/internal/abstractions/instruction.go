package abstractions

type Instruction struct {
	Operation    string
	LeftOperand  int64
	RightOperand int64
}

func (i Instruction) Execute() int64 {
	if i.Operation == "mul" {
		return i.LeftOperand * i.RightOperand
	}

	panic("Unknown operation")
}
