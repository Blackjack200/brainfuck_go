package brainfuck

import "github.com/dave/jennifer/jen"

func TransformToGo(p Program) string {
	if len(p) == 0 {
		f := jen.NewFile("main")
		f.Func().Id("main").Params().BlockFunc(func(g *jen.Group) {})
		return f.GoString()
	}
	p = transformOptimize(p)
	code := func() []Token {
		buf := make([]Token, len(p))
		for _, pp := range p {
			buf = append(buf, pp.Token)
		}
		return buf
	}()
	f := jen.NewFile("main")
	f.Func().Id("main").Params().BlockFunc(func(g *jen.Group) {
		g.Id("mem").Op(":=").Make(jen.Map(jen.Uint64()).Byte(), jen.Lit(30000))
		g.Id("ptr").Op(":=").Uint64().Op("(").Lit(0).Op(")")

		//TODO check usage
		if inArray(code, TokenInput) {
			g.Id("in").Op(":=").Func().Params().Block(
				jen.Var().Id("c").Byte(),
				jen.Qual("fmt", "Scanf").Call(jen.Lit("%c"), jen.Id("&c")),
				jen.Id("mem").Index(jen.Id("ptr")).Op("=").Id("c"),
			)
		}
		if inArray(code, TokenOutput) {
			g.Id("out").Op(":=").Func().Params().Block(
				jen.Qual("fmt", "Printf").Call(jen.Lit("%c"), jen.Id("mem").Index(jen.Id("ptr"))),
			)
		}

		for _, r := range p {
			switch r.Token {
			case TokenAdd:
				g.Id("mem").Index(jen.Id("ptr")).Op("+=").Lit(byte(r.Num))
			case TokenDec:
				g.Id("mem").Index(jen.Id("ptr")).Op("-=").Lit(byte(r.Num))
			case TokenMoveRight:
				g.Id("ptr").Op("+=").Lit(r.Num)
			case TokenMoveLeft:
				g.Id("ptr").Op("-=").Lit(r.Num)
			case TokenOutput:
				g.Id("out").Call()
			case TokenInput:
				g.Id("in").Call()
			case TokenStartLoop:
				g.For(jen.Id("mem").Index(jen.Id("ptr")).Op("!=").Lit(0)).Op("{")
			case TokenEndLoop:
				g.Op("}")
			}
		}
	})
	return f.GoString()
}
