package brainfuck

type SyntaxNode struct {
	Operator Token
	Parent   *SyntaxNode
	Children []*SyntaxNode
}

func Lex(t []Token) *SyntaxNode {
	root := &SyntaxNode{
		Operator: TokenProgramStart,
		Parent:   nil,
		Children: nil,
	}
	curt := root
	loopDepth := 0
	for _, token := range t {
		if token == TokenStartLoop {
			n := &SyntaxNode{
				Operator: TokenStartLoop,
				Parent:   curt,
			}
			curt.Children = append(curt.Children, n)
			curt = n
			loopDepth++
			continue
		}
		if token == TokenEndLoop {
			curt = curt.Parent
			loopDepth--
			continue
		}
		c := &SyntaxNode{
			Operator: token,
		}
		curt.Children = append(curt.Children, c)
	}
	return root
}
