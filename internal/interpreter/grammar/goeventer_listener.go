// Code generated from GoEventer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package grammar // GoEventer
import "github.com/antlr4-go/antlr/v4"

// GoEventerListener is a complete listener for a parse tree produced by GoEventerParser.
type GoEventerListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterStructStat is called when entering the structStat production.
	EnterStructStat(c *StructStatContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterAssignVariableStat is called when entering the assignVariableStat production.
	EnterAssignVariableStat(c *AssignVariableStatContext)

	// EnterUpdateVariableStat is called when entering the updateVariableStat production.
	EnterUpdateVariableStat(c *UpdateVariableStatContext)

	// EnterNotifyChanStat is called when entering the notifyChanStat production.
	EnterNotifyChanStat(c *NotifyChanStatContext)

	// EnterForStat is called when entering the forStat production.
	EnterForStat(c *ForStatContext)

	// EnterWhileStat is called when entering the whileStat production.
	EnterWhileStat(c *WhileStatContext)

	// EnterDefineListenerStat is called when entering the defineListenerStat production.
	EnterDefineListenerStat(c *DefineListenerStatContext)

	// EnterDefineFunctionStat is called when entering the defineFunctionStat production.
	EnterDefineFunctionStat(c *DefineFunctionStatContext)

	// EnterStatBlock is called when entering the statBlock production.
	EnterStatBlock(c *StatBlockContext)

	// EnterReturnStat is called when entering the returnStat production.
	EnterReturnStat(c *ReturnStatContext)

	// EnterMethodCallStat is called when entering the methodCallStat production.
	EnterMethodCallStat(c *MethodCallStatContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterMethodCallArguments is called when entering the methodCallArguments production.
	EnterMethodCallArguments(c *MethodCallArgumentsContext)

	// EnterFunctionDefineArguments is called when entering the functionDefineArguments production.
	EnterFunctionDefineArguments(c *FunctionDefineArgumentsContext)

	// EnterMethodCallExpr is called when entering the methodCallExpr production.
	EnterMethodCallExpr(c *MethodCallExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterMultiplicationExpr is called when entering the multiplicationExpr production.
	EnterMultiplicationExpr(c *MultiplicationExprContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterExprAtom is called when entering the exprAtom production.
	EnterExprAtom(c *ExprAtomContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// EnterStringAtom is called when entering the stringAtom production.
	EnterStringAtom(c *StringAtomContext)

	// EnterNilAtom is called when entering the nilAtom production.
	EnterNilAtom(c *NilAtomContext)

	// EnterMakeChanAtom is called when entering the makeChanAtom production.
	EnterMakeChanAtom(c *MakeChanAtomContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitStructStat is called when exiting the structStat production.
	ExitStructStat(c *StructStatContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitAssignVariableStat is called when exiting the assignVariableStat production.
	ExitAssignVariableStat(c *AssignVariableStatContext)

	// ExitUpdateVariableStat is called when exiting the updateVariableStat production.
	ExitUpdateVariableStat(c *UpdateVariableStatContext)

	// ExitNotifyChanStat is called when exiting the notifyChanStat production.
	ExitNotifyChanStat(c *NotifyChanStatContext)

	// ExitForStat is called when exiting the forStat production.
	ExitForStat(c *ForStatContext)

	// ExitWhileStat is called when exiting the whileStat production.
	ExitWhileStat(c *WhileStatContext)

	// ExitDefineListenerStat is called when exiting the defineListenerStat production.
	ExitDefineListenerStat(c *DefineListenerStatContext)

	// ExitDefineFunctionStat is called when exiting the defineFunctionStat production.
	ExitDefineFunctionStat(c *DefineFunctionStatContext)

	// ExitStatBlock is called when exiting the statBlock production.
	ExitStatBlock(c *StatBlockContext)

	// ExitReturnStat is called when exiting the returnStat production.
	ExitReturnStat(c *ReturnStatContext)

	// ExitMethodCallStat is called when exiting the methodCallStat production.
	ExitMethodCallStat(c *MethodCallStatContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitMethodCallArguments is called when exiting the methodCallArguments production.
	ExitMethodCallArguments(c *MethodCallArgumentsContext)

	// ExitFunctionDefineArguments is called when exiting the functionDefineArguments production.
	ExitFunctionDefineArguments(c *FunctionDefineArgumentsContext)

	// ExitMethodCallExpr is called when exiting the methodCallExpr production.
	ExitMethodCallExpr(c *MethodCallExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitMultiplicationExpr is called when exiting the multiplicationExpr production.
	ExitMultiplicationExpr(c *MultiplicationExprContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitExprAtom is called when exiting the exprAtom production.
	ExitExprAtom(c *ExprAtomContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)

	// ExitStringAtom is called when exiting the stringAtom production.
	ExitStringAtom(c *StringAtomContext)

	// ExitNilAtom is called when exiting the nilAtom production.
	ExitNilAtom(c *NilAtomContext)

	// ExitMakeChanAtom is called when exiting the makeChanAtom production.
	ExitMakeChanAtom(c *MakeChanAtomContext)
}
