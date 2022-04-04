package brainfuck

var optimizeToken = []byte{
	TokenAdd,
	TokenDec,
	TokenMoveRight,
	TokenMoveLeft,
}

func Optimize(old Program) Program {
	var newProgram []Opcode
	num := uint64(0)
	curtToken := Token(TokenProgramStart)

	mergeToken := func() {
		if num > 0 {
			//jmp not set
			newProgram = append(newProgram, Opcode{Token: curtToken, Num: num})
		}
		curtToken = Token(TokenProgramStart)
		num = 0
	}

	tokenType := map[Token]int{
		TokenAdd:       0,
		TokenDec:       0,
		TokenMoveRight: 1,
		TokenMoveLeft:  1,
	}

	for i, curt := range old {
		_ = old[i]
		if !inArray(optimizeToken, byte(curt.Token)) {
			mergeToken()
			newProgram = append(newProgram, curt)
			continue
		}
		if tokenType[curt.Token] == tokenType[curtToken] {
			if curt.Token != curtToken {
				if curt.Num > num {
					num = curt.Num - num
					curtToken = curt.Token
				} else {
					num -= curt.Num
				}
			} else {
				num += curt.Num
			}
		} else {
			mergeToken()
			curtToken = curt.Token
			num = curt.Num
		}
	}
	mergeToken()
	return fixIndex(newProgram)
}

func fixIndex(old Program) Program {
	var newProgram []Opcode
	var stack []int
	max := len(old)
	for i := 0; i < max; i++ {
		cur := old[i]
		if cur.Token == TokenStartLoop {
			stack = append(stack, len(newProgram))
			newProgram = append(newProgram, Opcode{
				Token: TokenStartLoop,
				Jmp:   i,
				Num:   cur.Num,
			})
		} else if cur.Token == TokenEndLoop {
			jmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			newProgram[jmp].Jmp = i
			newProgram = append(newProgram, Opcode{
				Token: TokenEndLoop,
				Jmp:   jmp,
				Num:   cur.Num,
			})
		} else {
			newProgram = append(newProgram, cur)
		}
	}
	return newProgram
}
