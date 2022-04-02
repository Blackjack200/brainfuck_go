package main

import (
	"brainfuck/brainfuck"
	"brainfuck/brainfuck/bfruntime"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, err := ioutil.ReadFile("test.bf")
	if err != nil {
		panic(fmt.Errorf("error reading file: %v", err))
	}
	tokens, err := brainfuck.Parse(buf)
	if err != nil {
		panic(fmt.Errorf("error lexing file: %v", err))
	}
	m := bfruntime.NewGoMachine(os.Stdin, os.Stdout)
	nodes := brainfuck.Lex(tokens)
	program := brainfuck.Compile(nodes)
	//newProgram := brainfuck.Optimize(program)
	//newProgram.Run(m)
	program.Run(m)
}
