package brainfuck

import (
	"brainfuck/brainfuck/bfruntime"
	"fmt"
)

type Opcode struct {
	Token Token
	Jmp   int
	Num   uint64
}

type Program []Opcode

func (p Program) Run(m bfruntime.Machine) {
	for pos, max := 0, len(p); pos < max; pos++ {
		c := p[pos]
		jmp := func() {
			pos = c.Jmp
		}
		switch c.Token {
		case TokenAdd:
			m.Tape().Inc(c.Num)
		case TokenDec:
			m.Tape().Dec(c.Num)
		case TokenOutput:
			m.Out(m.Tape().Read())
		case TokenInput:
			m.Tape().Write(m.In())
		case TokenMoveRight:
			m.Tape().Next(c.Num)
		case TokenMoveLeft:
			m.Tape().Prev(c.Num)
		case TokenStartLoop:
			if m.Tape().Read() == 0 {
				jmp()
			}
		case TokenEndLoop:
			if m.Tape().Read() != 0 {
				jmp()
			}
		}
	}
}

func toOpcode(root []*SyntaxNode, ptr *[]Opcode) {
	for i, m := 0, len(root); i < m; i++ {
		r := root[i]
		switch r.Operator {
		case TokenStartLoop:
			startLoopPos := len(*ptr)
			*ptr = append(*ptr, Opcode{Token: r.Operator, Num: 1})
			toOpcode(r.Children, ptr)
			endLoopPos := len(*ptr)
			*ptr = append(*ptr, Opcode{Token: TokenEndLoop, Jmp: startLoopPos, Num: 1})
			(*ptr)[startLoopPos].Jmp = endLoopPos
		case TokenProgramStart:
			if len(*ptr) != 0 {
				panic(fmt.Errorf("invalid program start"))
			}
		case TokenAdd,
			TokenDec,
			TokenOutput,
			TokenInput,
			TokenMoveRight,
			TokenMoveLeft:
			*ptr = append(*ptr, Opcode{Token: r.Operator, Num: 1})
		}
	}
}

func Compile(root *SyntaxNode) Program {
	var program []Opcode
	toOpcode(root.Children, &program)
	return program
}
