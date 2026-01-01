package abstractions

type Program struct {
	instructions []Instruction
}

func NewProgram(
	instructions []Instruction,
) *Program {
	return &Program{instructions}
}

func (p *Program) GetInstructionsCount() uint {
	return uint(len(p.instructions))
}

func (p *Program) Execute() int64 {
	total := int64(0)

	for _, instruction := range p.instructions {
		total += instruction.Execute()
	}

	return total
}
