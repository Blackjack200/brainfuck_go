package brainfuck

var optimizeToken = []byte{
	TokenAdd,
	TokenDec,
	TokenMoveRight,
	TokenMoveLeft,
}

func transformOptimize(old Program) Program {
	//TODO more optimization
	var newProgram []Opcode
	num := uint64(0)
	curtToken := Token(TokenProgramStart)

	for pos, max := 0, len(old); pos < max; pos++ {
		cur := old[pos]
		mergeToken := func() {
			if num > 0 {
				//jmp not set
				newProgram = append(newProgram, Opcode{Token: curtToken, Num: num})
			}
			curtToken = Token(TokenProgramStart)
		}
		if !inArray(optimizeToken, byte(cur.Token)) {
			mergeToken()
			num = 0
			newProgram = append(newProgram, cur)
			continue
		}
		if cur.Token == curtToken {
			num++
		} else {
			mergeToken()
			curtToken = cur.Token
			num = 1
		}
	}
	newProgram = append(newProgram, Opcode{Token: curtToken, Num: num})
	return newProgram
}
