// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by ExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by ExprParser#ip.
	VisitIp(ctx *IpContext) interface{}

	// Visit a parse tree produced by ExprParser#impact.
	VisitImpact(ctx *ImpactContext) interface{}

	// Visit a parse tree produced by ExprParser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by ExprParser#listIn.
	VisitListIn(ctx *ListInContext) interface{}

	// Visit a parse tree produced by ExprParser#addSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by ExprParser#mulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by ExprParser#re.
	VisitRe(ctx *ReContext) interface{}

	// Visit a parse tree produced by ExprParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by ExprParser#reg.
	VisitReg(ctx *RegContext) interface{}

	// Visit a parse tree produced by ExprParser#cidr.
	VisitCidr(ctx *CidrContext) interface{}

	// Visit a parse tree produced by ExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ExprParser#logic.
	VisitLogic(ctx *LogicContext) interface{}

	// Visit a parse tree produced by ExprParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}
}
