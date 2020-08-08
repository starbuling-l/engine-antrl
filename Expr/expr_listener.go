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

	// EnterIp is called when entering the ip production.
	EnterIp(c *IpContext)

	// EnterImpact is called when entering the impact production.
	EnterImpact(c *ImpactContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterListIn is called when entering the listIn production.
	EnterListIn(c *ListInContext)

	// EnterAddSub is called when entering the addSub production.
	EnterAddSub(c *AddSubContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterMulDiv is called when entering the mulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterRe is called when entering the re production.
	EnterRe(c *ReContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterReg is called when entering the reg production.
	EnterReg(c *RegContext)

	// EnterCidr is called when entering the cidr production.
	EnterCidr(c *CidrContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterLogic is called when entering the logic production.
	EnterLogic(c *LogicContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitIp is called when exiting the ip production.
	ExitIp(c *IpContext)

	// ExitImpact is called when exiting the impact production.
	ExitImpact(c *ImpactContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitListIn is called when exiting the listIn production.
	ExitListIn(c *ListInContext)

	// ExitAddSub is called when exiting the addSub production.
	ExitAddSub(c *AddSubContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitMulDiv is called when exiting the mulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitRe is called when exiting the re production.
	ExitRe(c *ReContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitReg is called when exiting the reg production.
	ExitReg(c *RegContext)

	// ExitCidr is called when exiting the cidr production.
	ExitCidr(c *CidrContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitLogic is called when exiting the logic production.
	ExitLogic(c *LogicContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)
}
