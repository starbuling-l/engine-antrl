// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIP is called when entering the IP production.
	EnterIP(c *IPContext)

	// EnterImpact is called when entering the impact production.
	EnterImpact(c *ImpactContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterAddSub is called when entering the addSub production.
	EnterAddSub(c *AddSubContext)

	// EnterListIn is called when entering the listIn production.
	EnterListIn(c *ListInContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterLogic is called when entering the logic production.
	EnterLogic(c *LogicContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterMulDiv is called when entering the mulDiv production.
	EnterMulDiv(c *MulDivContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIP is called when exiting the IP production.
	ExitIP(c *IPContext)

	// ExitImpact is called when exiting the impact production.
	ExitImpact(c *ImpactContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitAddSub is called when exiting the addSub production.
	ExitAddSub(c *AddSubContext)

	// ExitListIn is called when exiting the listIn production.
	ExitListIn(c *ListInContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitLogic is called when exiting the logic production.
	ExitLogic(c *LogicContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitMulDiv is called when exiting the mulDiv production.
	ExitMulDiv(c *MulDivContext)
}
