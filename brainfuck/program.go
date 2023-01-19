package brainfuck

import (
	"brainfuck/brainfuck/bfruntime"
)

type Program []Opcode

func (p Program) Run(m bfruntime.Machine) {
	for pos, max := 0, len(p); pos < max; pos++ {
		c := p[pos]
		jmp := func() {
			pos = c.Jmp
		}
		switch c.Token {
		case TokenAdd:
			p.dumpMem(m)
			m.Tape().Inc(c.Num)
		case TokenDec:
			p.dumpMem(m)
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

func (p Program) dumpMem(m bfruntime.Machine) {
	m.Tape().Dump()
}
