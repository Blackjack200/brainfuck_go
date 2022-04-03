package main

import (
	"brainfuck/brainfuck"
	"fmt"
	"io/ioutil"
)

func main() {
	writeFile := func(file string, data string) {
		err := ioutil.WriteFile(file, []byte(data), 0644)
		if err != nil {
			panic(err)
		}
	}
	buf, err := ioutil.ReadFile("test.bf")
	if err != nil {
		panic(fmt.Errorf("error reading file: %v", err))
	}
	tokens, err := brainfuck.Parse(buf)
	if err != nil {
		panic(fmt.Errorf("error lexing file: %v", err))
	}
	//m := bfruntime.NewGoMachine(os.Stdin, os.Stdout)
	nodes := brainfuck.Lex(tokens)
	program := brainfuck.Compile(nodes)
	writeFile("test_opt.go", brainfuck.TransformToGo(program))
	//newProgram := brainfuck.Optimize(program)
	//newProgram.Run(m)
	//program.Run(m)
}
