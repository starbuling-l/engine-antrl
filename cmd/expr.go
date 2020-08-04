package main

import (
	parser "RuleCalculationEngine/antlrdemo/Expr"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strconv"
	"strings"
)


type EvalVisitor struct {
	parser.ExprVisitor
	memory map[string]interface{}
	//list []string
}

func NewEvaVisitor()*EvalVisitor {
	visitor := new(EvalVisitor)
	visitor.memory = make(map[string]interface{})
	//visitor.list = make([]string,10)
	return visitor
}

/*
*func (prc *BaseParserRuleContext) Accept(visitor ParseTreeVisitor) interface{} {
	return visitor.VisitChildren(prc)
}
*/

func (this *EvalVisitor)Visit(node antlr.ParseTree) interface{} {
	res := node.Accept(this)
	return res
}
/*func (this *EvalVisitor)VisitProgram(ctx *parser.ProgramContext) interface{} {
	value := this.Visit(ctx.GetRuleContext())
	return value
}*/

func (this *EvalVisitor)VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	value := this.Visit(ctx.Expr())
	return value
}

func (this *EvalVisitor)VisitAssign(ctx *parser.AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := this.Visit(ctx.Expr())
	this.memory[id] = value
	return value
}

func (this *EvalVisitor)VisitParens(ctx *parser.ParensContext) interface{} {
	return this.Visit(ctx.Expr())
}

func (this *EvalVisitor)VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	value, ok := this.memory[id]
	if ok {
		return value
	}
	return nil
}

func (this *EvalVisitor) VisitList(ctx *parser.ListContext) interface{} {
	//value := this.Visit(ctx.List())
	list := make([]string,10)
	value := strings.Trim(ctx.List().GetText(), "[")
	value = strings.Trim(value, "]")
	list = strings.Split(value,",")
	return list
}

func (this *EvalVisitor)VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	value :=ctx.StringLiteral().GetText()
	return value
}

func (this *EvalVisitor)VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	value, _ := strconv.Atoi( ctx.IntLiteral().GetText())
	return value
}

func (this *EvalVisitor)VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	value, _ := strconv.ParseBool( ctx.BooleanLiteral().GetText())
	return value
}

func (this *EvalVisitor)VisitIP(ctx *parser.IPContext) interface{} {
	value:=ctx.IP().GetText()
	return value
}

func (this *EvalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(int)
	right := this.Visit(ctx.Expr(1)).(int)
	if ctx.GetOp().GetTokenType() == parser.ExprLexerADD {
		return left + right
	}
	return left - right
}

func (this *EvalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(int)
	right := this.Visit(ctx.Expr(1)).(int)
	if ctx.GetOp().GetTokenType() == parser.ExprLexerMUL {
		return left * right
	}
	return left / right
}

func (this *EvalVisitor) VisitLogic(ctx *parser.LogicContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(int)
	right := this.Visit(ctx.Expr(1)).(int)
	flag := ctx.GetOp().GetTokenType()
	switch flag{
	case parser.ExprLexerGT:
		return left > right
	case parser.ExprLexerLT:
		return left < right
	case parser.ExprLexerEQUAL:
		return left == right
	case parser.ExprParserLE:
		return left <= right
	case parser.ExprLexerGE:
		return left >=right
	}
	return nil
}

func (this *EvalVisitor) VisitImpact(ctx *parser.ImpactContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(bool)
	right := this.Visit(ctx.Expr(1)).(bool)
	if ctx.GetOp().GetTokenType() == parser.ExprLexerAND {
		return left && right
	}
	return left || right
}

func (this *EvalVisitor) VisitListIn(ctx *parser.ListInContext) interface{} {
	left := strconv.Itoa(this.Visit(ctx.Expr(0)).(int))
	right := this.Visit(ctx.Expr(1)).([]string)
	fmt.Println(right)
	for _,value:=range right{
		fmt.Println(value)
		if value == left{
			return true
		}
	}
	return false
}

//func Calc(input string) interface{} {
//	inputStream := antlr.NewInputStream(input)
//	lexer := parser.NewExprLexer(inputStream)
//	tokens := antlr.NewCommonTokenStream(lexer, 0)
//	parse := parser.NewExprParser(tokens)
//
//	tree := parse.Stat()
//	//fmt.Println(tree.GetText())
//	return tree.Accept(NewEvaVisitor())
//}

func main()  {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewExprLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parse := parser.NewExprParser(tokens)

	tree := parse.Stat()
	fmt.Println(tree.Accept(NewEvaVisitor()))
	//fmt.Printf("d = 2     : %d\n", Calc("d = 2"))
	//fmt.Printf("3         : %d\n", Calc("3"))
	//fmt.Printf("3 * 4     : %d\n", Calc("3 * 4"))
	//fmt.Printf("b = 5 * 6 : %d\n", Calc("b = 5 * 6"))
	//fmt.Printf("2 / 2     : %d\n", Calc("2 / 2"))
	//fmt.Printf("c = 4 / 2 : %d\n", Calc("c = 4 / 2"))
	//fmt.Printf("4 / 2 + 3 : %d\n", Calc("4 / 2 + 3"))
	//fmt.Printf("3 + 4 / 2 : %d\n", Calc("3 + 4 / 2"))
	//fmt.Printf("4 / 2 * 2 : %d\n", Calc("4 / 2 * 2"))
	//fmt.Println(Calc("age = 1 + 1 +4 / (2 * 2)"))
	//fmt.Println(Calc(" 1 < 2 "))
	//fmt.Println(Calc("( 1 < 2 ) && ( 1>2 )"))
	//fmt.Println(Calc("false || true"))
	//fmt.Println(Calc("true && true"))
	//fmt.Println(Calc("2 in [1,2,3]"))
	//fmt.Println(Calc("5 in [1,2,3]"))
	//fmt.Println(Calc("127.0.0.1"))
}





