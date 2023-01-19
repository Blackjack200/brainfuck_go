package main

import (
	"brainfuck/brainfuck"
	"brainfuck/brainfuck/bfruntime"
	"fmt"
	"os"
	"time"
)

func main() {
	writeFile := func(file string, data string) {
		err := os.WriteFile(file, []byte(data), 0644)
		if err != nil {
			panic(err)
		}
	}
	buf, err := os.ReadFile("test.bf")
	if err != nil {
		panic(fmt.Errorf("error reading file: %v", err))
	}
	tokens, err := brainfuck.Parse(buf)
	if err != nil {
		panic(fmt.Errorf("error lexing file: %v", err))
	}
	m := bfruntime.NewGoMachine(bfruntime.NewCTape(), os.Stdin, os.Stdout)
	nodes := brainfuck.Lex(tokens)
	program := brainfuck.Compile(nodes)
	newProgram := brainfuck.Optimize(program)
	writeFile("test_opt.go", brainfuck.TransformToGo(newProgram))
	start := time.Now()
	program.Run(m)
	fmt.Printf("time used: %v\n", time.Now().Sub(start).String())
	//program.Run(m)
}
