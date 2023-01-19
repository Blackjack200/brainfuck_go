package bfruntime

import (
	"fmt"
	"io"
)

type Tape interface {
	Next(uint64)
	Prev(uint64)
	Write(byte)
	Read() byte
	Inc(uint64)
	Dec(uint64)
	Dump() map[uint64]byte
}

type Machine interface {
	Tape() Tape
	In() byte
	Out(byte)
}

type GoMachine struct {
	tape Tape
	in   io.Reader
	out  io.Writer
}

func (g *GoMachine) Tape() Tape {
	return g.tape
}

func (g *GoMachine) In() byte {
	var c byte
	_, _ = fmt.Fscanf(g.in, "%c", &c)
	return c
}

func (g *GoMachine) Out(b byte) {
	_, _ = fmt.Fprintf(g.out, "%c", b)
}

func NewGoMachine(tape Tape, in io.Reader, out io.Writer) *GoMachine {
	return &GoMachine{
		tape: tape,
		in:   in,
		out:  out,
	}
}
