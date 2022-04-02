package brainfuck

var optimizeToken = []byte{
	TokenAdd,
	TokenDec,
	TokenMoveRight,
	TokenMoveLeft,
	TokenOutput,
}
/*
func Optimize(old Program) Program {
	var newProgram []Opcode
	num := uint64(0)
	curtToken := Token(TokenProgramStart)
	for pos, max := 0, len(old); pos < max; pos++ {
		cur := old[pos]
		if !inArray(optimizeToken, byte(cur.Token)) {
			newProgram = append(newProgram, cur)
			continue
		}
		if cur.Token == curtToken {
			num++
		} else {
			if num > 0 {
				//jmp not set
				newProgram = append(newProgram, Opcode{Token: curtToken, Num: num})
			}
			num = 1
			curtToken = cur.Token
		}
	}
	newProgram = append(newProgram, Opcode{Token: curtToken, Num: num})
	return newProgram
}*/