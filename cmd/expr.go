package Expr

import (
	parser "RuleEngine/Expr"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"regexp"
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

//func (this *EvalVisitor)VisitId(ctx *parser.IdContext) interface{} {
//	id := ctx.ID().GetText()
//	value, ok := this.memory[id]
//	if ok {
//		return value
//	}
//	return nil
//}

func (this *EvalVisitor)VisitId(ctx *parser.IdContext) interface{} {
	value := ctx.ID().GetText()
	return value
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
func (this *EvalVisitor)VisitReg(ctx *parser.RegContext) interface{} {
	value := strings.Trim(ctx.Reg().GetText(), "`")
	fmt.Println(value)
	return value
}
func (this *EvalVisitor)VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	value, _ := strconv.ParseBool( ctx.BooleanLiteral().GetText())
	return value
}

func (this *EvalVisitor)VisitIp(ctx *parser.IpContext) interface{} {
	value:=ctx.IP().GetText()
	return value
}

func (this *EvalVisitor)VisitCidr(ctx *parser.CidrContext) interface{} {
	value:= ctx.CIDR().GetText()
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

func (this *EvalVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {

	left := this.ToString(ctx.Expr(0))
	right := this.ToString(ctx.Expr(1))
	if ctx.GetOp().GetTokenType() == parser.ExprLexerNOTEQUAL{
		return left != right
	}
	return left == right
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
	case parser.ExprParserLE:
		return left <= right
	case parser.ExprLexerGE:
		return left >=right
	}
	return false
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
	left := this.Visit(ctx.Expr(0)).(string)
	fmt.Println(left)
	right := this.Visit(ctx.Expr(1)).([]string)
	fmt.Println(right)
	//fmt.Println(right)
	for _,value:=range right{
		//fmt.Println(value)
		if value == left{
			return true
		}
	}
	return false
}

func (this *EvalVisitor) VisitRe(ctx *parser.ReContext) interface{} {
	left :=this.ToString(this.Visit(ctx.Expr(0)))
	fmt.Println(left)
	right := this.Visit(ctx.Expr(1)).(string)
	fmt.Println(right)
	if ok, _ := regexp.Match(right, []byte(left)); ok {
		return true
	}
	return false
}

func(this *EvalVisitor) ToString(visitNode interface{}) (ret string){
	switch visitNode.(type){
	case string:
		ret = visitNode.(string)
	case int:
		ret = strconv.Itoa(visitNode.(int))
	}
	return
}

func Calc(rule string,context map[string]string) interface{} {
	for k,v:=range context{
		rule =strings.ReplaceAll(rule,"$"+k,v)
	}
	println(rule)
	inputStream := antlr.NewInputStream(rule)
	lexer := parser.NewExprLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parse := parser.NewExprParser(tokens)

	tree := parse.Stat()
	//fmt.Println(tree.GetText())
	return tree.Accept(NewEvaVisitor())
}

//func main()  {
//	input, _ := antlr.NewFileStream(os.Args[1])
//	lexer := parser.NewExprLexer(input)
//	tokens := antlr.NewCommonTokenStream(lexer, 0)
//	parse := parser.NewExprParser(tokens)
//
//	tree := parse.Stat()
//	fmt.Println(tree.Accept(NewEvaVisitor()))
//	//fmt.Printf("d = 2     : %d\n", Calc("d = 2"))
//	//fmt.Printf("3         : %d\n", Calc("3"))
//	//fmt.Printf("3 * 4     : %d\n", Calc("3 * 4"))
//	//fmt.Printf("_b = 5 * 6 : %d\n", Calc("_b = 5 * 6"))
//	//fmt.Printf("2 / 2     : %d\n", Calc("2 / 2"))
//	//fmt.Printf("c = 4 / 2 : %d\n", Calc("c = 4 / 2"))
//	//fmt.Printf("4 / 2 + 3 : %d\n", Calc("4 / 2 + 3"))
//	//fmt.Printf("3 + 4 / 2 : %d\n", Calc("3 + 4 / 2"))
//	//fmt.Printf("4 / 2 * 2 : %d\n", Calc("4 / 2 * 2"))
//	//fmt.Printf("age = 1 + 1 +4 / (2 * 2) : %d\n",Calc("age = 1 + 1 +4 / (2 * 2)"))
//	//fmt.Printf("1 < 2  : %t\n",Calc(" 1 < 2 "))
//	//fmt.Printf("( 1 < 2 ) && ( 1>2 ) : %t\n",Calc("( 1 < 2 ) && ( 1>2 )"))
//	//fmt.Printf("false || true : %t\n",Calc("false || true"))
//	//fmt.Printf("true && true : %t\n",Calc("true && true"))
//	//fmt.Printf("2 in [1,2,3] : %t\n",Calc("2 in [1,2,3]"))
//	//fmt.Printf("5 in [1,2,3] : %t\n",Calc("5 in [1,2,3]"))
//	//fmt.Printf("127.0.0.1 : %v\n",Calc("127.0.0.1"))
//	//fmt.Printf("127.0.0.1/6 : %v\n",Calc("127.0.0.1/6"))
//	//fmt.Printf("\"123s\" re \"[0-9]+\" : %t\n",Calc("\"123s\" re \"[0-9]+\""))
//	//fmt.Printf("\"123\" re \"[0-9]+\" : %t\n",Calc("\"123\" re \"[0-9]+\""))
//}





