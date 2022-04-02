package brainfuck

import "fmt"

type Token byte

//TokenProgramStart internal token
const (
	TokenProgramStart = 'S'
	TokenStartLoop    = '['
	TokenEndLoop      = ']'
	TokenAdd          = '+'
	TokenDec          = '-'
	TokenOutput       = '.'
	TokenInput        = ','
	TokenMoveRight    = '>'
	TokenMoveLeft     = '<'
)

var validToken = []byte{'+', '-', '.', ',', '>', '<', '[', ']', '~'}
var commentToken = []byte{';', '#'}

func Parse(buf []byte) ([]Token, error) {
	line := 0
	pos := 0
	ignored := false
	loopCounter := 0
	var tokens []Token
	for _, tok := range buf {
		pos++
		if tok == '\n' {
			line++
			pos = 0
			ignored = false
			continue
		}
		if ignored {
			continue
		}
		if inArray(commentToken, tok) {
			ignored = true
			continue
		}
		if tok == TokenStartLoop {
			loopCounter++
		}
		if tok == TokenEndLoop {
			loopCounter--
			if loopCounter < 0 {
				return nil, fmt.Errorf("unmatched ']' at line %d, pos %d", line, pos)
			}
		}
		if inArray(validToken, tok) {
			tokens = append(tokens, Token(tok))
			continue
		}
	}
	if loopCounter != 0 {
		return nil, fmt.Errorf("unmatched '[' at line %d, pos %d", line, pos)
	}
	return tokens, nil
}
