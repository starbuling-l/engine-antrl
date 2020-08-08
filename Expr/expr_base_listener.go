// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseExprListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseExprListener) ExitProgram(ctx *ProgramContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseExprListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseExprListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseExprListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseExprListener) ExitAssign(ctx *AssignContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseExprListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseExprListener) ExitParens(ctx *ParensContext) {}

// EnterIp is called when production ip is entered.
func (s *BaseExprListener) EnterIp(ctx *IpContext) {}

// ExitIp is called when production ip is exited.
func (s *BaseExprListener) ExitIp(ctx *IpContext) {}

// EnterImpact is called when production impact is entered.
func (s *BaseExprListener) EnterImpact(ctx *ImpactContext) {}

// ExitImpact is called when production impact is exited.
func (s *BaseExprListener) ExitImpact(ctx *ImpactContext) {}

// EnterIntLiteral is called when production intLiteral is entered.
func (s *BaseExprListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production intLiteral is exited.
func (s *BaseExprListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterListIn is called when production listIn is entered.
func (s *BaseExprListener) EnterListIn(ctx *ListInContext) {}

// ExitListIn is called when production listIn is exited.
func (s *BaseExprListener) ExitListIn(ctx *ListInContext) {}

// EnterAddSub is called when production addSub is entered.
func (s *BaseExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production addSub is exited.
func (s *BaseExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterList is called when production list is entered.
func (s *BaseExprListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseExprListener) ExitList(ctx *ListContext) {}

// EnterMulDiv is called when production mulDiv is entered.
func (s *BaseExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production mulDiv is exited.
func (s *BaseExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseExprListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseExprListener) ExitEqual(ctx *EqualContext) {}

// EnterRe is called when production re is entered.
func (s *BaseExprListener) EnterRe(ctx *ReContext) {}

// ExitRe is called when production re is exited.
func (s *BaseExprListener) ExitRe(ctx *ReContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseExprListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseExprListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterReg is called when production reg is entered.
func (s *BaseExprListener) EnterReg(ctx *RegContext) {}

// ExitReg is called when production reg is exited.
func (s *BaseExprListener) ExitReg(ctx *RegContext) {}

// EnterCidr is called when production cidr is entered.
func (s *BaseExprListener) EnterCidr(ctx *CidrContext) {}

// ExitCidr is called when production cidr is exited.
func (s *BaseExprListener) ExitCidr(ctx *CidrContext) {}

// EnterId is called when production id is entered.
func (s *BaseExprListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseExprListener) ExitId(ctx *IdContext) {}

// EnterLogic is called when production logic is entered.
func (s *BaseExprListener) EnterLogic(ctx *LogicContext) {}

// ExitLogic is called when production logic is exited.
func (s *BaseExprListener) ExitLogic(ctx *LogicContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseExprListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseExprListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
