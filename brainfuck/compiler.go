package brainfuck

import (
	"fmt"
)

type Opcode struct {
	Token Token
	Jmp   int
	Num   uint64
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
