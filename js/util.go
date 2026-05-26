package js

import (
	"github.com/tdewolff/parse/v2/js"
)

var (
	spaceBytes                 = []byte(" ")
	newlineBytes               = []byte("\n")
	starBytes                  = []byte("*")
	plusBytes                  = []byte("+")
	plusPlusBytes              = []byte("++")
	minMinBytes                = []byte("--")
	expBytes                   = []byte("**")
	bitOrBytes                 = []byte("|")
	colonBytes                 = []byte(":")
	semicolonBytes             = []byte(";")
	commaBytes                 = []byte(",")
	dotBytes                   = []byte(".")
	ellipsisBytes              = []byte("...")
	openBraceBytes             = []byte("{")
	closeBraceBytes            = []byte("}")
	openParenBytes             = []byte("(")
	closeParenBytes            = []byte(")")
	openBracketBytes           = []byte("[")
	closeBracketBytes          = []byte("]")
	openParenBracketBytes      = []byte("({")
	closeParenOpenBracketBytes = []byte("){")
	notBytes                   = []byte("!")
	questionBytes              = []byte("?")
	equalBytes                 = []byte("=")
	optChainBytes              = []byte("?.")
	arrowBytes                 = []byte("=>")
	notEqualBytes              = []byte("!=")
	zeroBytes                  = []byte("0")
	oneBytes                   = []byte("1")
	letBytes                   = []byte("let")
	getBytes                   = []byte("get")
	setBytes                   = []byte("set")
	asyncBytes                 = []byte("async")
	functionBytes              = []byte("function")
	staticBytes                = []byte("static")
	ifOpenBytes                = []byte("if(")
	elseBytes                  = []byte("else")
	withOpenBytes              = []byte("with(")
	doBytes                    = []byte("do")
	whileOpenBytes             = []byte("while(")
	forOpenBytes               = []byte("for(")
	forAwaitOpenBytes          = []byte("for await(")
	inBytes                    = []byte("in")
	ofBytes                    = []byte("of")
	switchOpenBytes            = []byte("switch(")
	throwBytes                 = []byte("throw")
	tryBytes                   = []byte("try")
	catchBytes                 = []byte("catch")
	finallyBytes               = []byte("finally")
	importBytes                = []byte("import")
	exportBytes                = []byte("export")
	fromBytes                  = []byte("from")
	returnBytes                = []byte("return")
	classBytes                 = []byte("class")
	asSpaceBytes               = []byte("as ")
	asyncSpaceBytes            = []byte("async ")
	spaceDefaultBytes          = []byte(" default")
	spaceExtendsBytes          = []byte(" extends")
	yieldBytes                 = []byte("yield")
	newBytes                   = []byte("new")
	openNewBytes               = []byte("(new")
	newTargetBytes             = []byte("new.target")
	importMetaBytes            = []byte("import.meta")
	nanBytes                   = []byte("NaN")
	undefinedBytes             = []byte("undefined")
	infinityBytes              = []byte("Infinity")
	nullBytes                  = []byte("null")
	zeroIndexBytes             = []byte("0[0]")
	groupedZeroIndexBytes      = []byte("(0[0])")
	oneDivZeroBytes            = []byte("1/0")
	groupedOneDivZeroBytes     = []byte("(1/0)")
	notZeroBytes               = []byte("!0")
	groupedNotZeroBytes        = []byte("(!0)")
	notOneBytes                = []byte("!1")
	groupedNotOneBytes         = []byte("(!1)")
	debuggerBytes              = []byte("debugger")
	regExpScriptBytes          = []byte("/script>")
	isNaNBytes                 = []byte("isNaN")
	NumberBytes                = []byte("Number")
	MathBytes                  = []byte("Math")
)

func isEmptyStmt(stmt js.IStmt) bool { _ = "STUB: not implemented"; return false }

func isFlowStmt(stmt js.IStmt) bool { _ = "STUB: not implemented"; return false }

func lastStmt(stmt js.IStmt) js.IStmt { _ = "STUB: not implemented"; return *new(js.IStmt) }

func endsInIf(istmt js.IStmt) bool { _ = "STUB: not implemented"; return false }

// precedence maps for the precedence inside the operation
var unaryPrecMap = map[js.TokenType]js.OpPrec{
	js.PostIncrToken: js.OpLHS,
	js.PostDecrToken: js.OpLHS,
	js.PreIncrToken:  js.OpUnary,
	js.PreDecrToken:  js.OpUnary,
	js.NotToken:      js.OpUnary,
	js.BitNotToken:   js.OpUnary,
	js.TypeofToken:   js.OpUnary,
	js.VoidToken:     js.OpUnary,
	js.DeleteToken:   js.OpUnary,
	js.PosToken:      js.OpUnary,
	js.NegToken:      js.OpUnary,
	js.AwaitToken:    js.OpUnary,
}

var binaryLeftPrecMap = map[js.TokenType]js.OpPrec{
	js.EqToken:         js.OpLHS,
	js.MulEqToken:      js.OpLHS,
	js.DivEqToken:      js.OpLHS,
	js.ModEqToken:      js.OpLHS,
	js.ExpEqToken:      js.OpLHS,
	js.AddEqToken:      js.OpLHS,
	js.SubEqToken:      js.OpLHS,
	js.LtLtEqToken:     js.OpLHS,
	js.GtGtEqToken:     js.OpLHS,
	js.GtGtGtEqToken:   js.OpLHS,
	js.BitAndEqToken:   js.OpLHS,
	js.BitXorEqToken:   js.OpLHS,
	js.BitOrEqToken:    js.OpLHS,
	js.ExpToken:        js.OpUpdate,
	js.MulToken:        js.OpMul,
	js.DivToken:        js.OpMul,
	js.ModToken:        js.OpMul,
	js.AddToken:        js.OpAdd,
	js.SubToken:        js.OpAdd,
	js.LtLtToken:       js.OpShift,
	js.GtGtToken:       js.OpShift,
	js.GtGtGtToken:     js.OpShift,
	js.LtToken:         js.OpCompare,
	js.LtEqToken:       js.OpCompare,
	js.GtToken:         js.OpCompare,
	js.GtEqToken:       js.OpCompare,
	js.InToken:         js.OpCompare,
	js.InstanceofToken: js.OpCompare,
	js.EqEqToken:       js.OpEquals,
	js.NotEqToken:      js.OpEquals,
	js.EqEqEqToken:     js.OpEquals,
	js.NotEqEqToken:    js.OpEquals,
	js.BitAndToken:     js.OpBitAnd,
	js.BitXorToken:     js.OpBitXor,
	js.BitOrToken:      js.OpBitOr,
	js.AndToken:        js.OpAnd,
	js.OrToken:         js.OpOr,
	js.NullishToken:    js.OpBitOr, // or OpCoalesce
	js.CommaToken:      js.OpExpr,
}

var binaryRightPrecMap = map[js.TokenType]js.OpPrec{
	js.EqToken:         js.OpAssign,
	js.MulEqToken:      js.OpAssign,
	js.DivEqToken:      js.OpAssign,
	js.ModEqToken:      js.OpAssign,
	js.ExpEqToken:      js.OpAssign,
	js.AddEqToken:      js.OpAssign,
	js.SubEqToken:      js.OpAssign,
	js.LtLtEqToken:     js.OpAssign,
	js.GtGtEqToken:     js.OpAssign,
	js.GtGtGtEqToken:   js.OpAssign,
	js.BitAndEqToken:   js.OpAssign,
	js.BitXorEqToken:   js.OpAssign,
	js.BitOrEqToken:    js.OpAssign,
	js.ExpToken:        js.OpExp,
	js.MulToken:        js.OpExp,
	js.DivToken:        js.OpExp,
	js.ModToken:        js.OpExp,
	js.AddToken:        js.OpMul,
	js.SubToken:        js.OpMul,
	js.LtLtToken:       js.OpAdd,
	js.GtGtToken:       js.OpAdd,
	js.GtGtGtToken:     js.OpAdd,
	js.LtToken:         js.OpShift,
	js.LtEqToken:       js.OpShift,
	js.GtToken:         js.OpShift,
	js.GtEqToken:       js.OpShift,
	js.InToken:         js.OpShift,
	js.InstanceofToken: js.OpShift,
	js.EqEqToken:       js.OpCompare,
	js.NotEqToken:      js.OpCompare,
	js.EqEqEqToken:     js.OpCompare,
	js.NotEqEqToken:    js.OpCompare,
	js.BitAndToken:     js.OpEquals,
	js.BitXorToken:     js.OpBitAnd,
	js.BitOrToken:      js.OpBitXor,
	js.AndToken:        js.OpAnd,   // changes order in AST but not in execution
	js.OrToken:         js.OpOr,    // changes order in AST but not in execution
	js.NullishToken:    js.OpBitOr, // or OpCoalesce
	js.CommaToken:      js.OpAssign,
}

// precedence maps of the operation itself
var unaryOpPrecMap = map[js.TokenType]js.OpPrec{
	js.PostIncrToken: js.OpUpdate,
	js.PostDecrToken: js.OpUpdate,
	js.PreIncrToken:  js.OpUpdate,
	js.PreDecrToken:  js.OpUpdate,
	js.NotToken:      js.OpUnary,
	js.BitNotToken:   js.OpUnary,
	js.TypeofToken:   js.OpUnary,
	js.VoidToken:     js.OpUnary,
	js.DeleteToken:   js.OpUnary,
	js.PosToken:      js.OpUnary,
	js.NegToken:      js.OpUnary,
	js.AwaitToken:    js.OpUnary,
}

var binaryOpPrecMap = map[js.TokenType]js.OpPrec{
	js.EqToken:         js.OpAssign,
	js.MulEqToken:      js.OpAssign,
	js.DivEqToken:      js.OpAssign,
	js.ModEqToken:      js.OpAssign,
	js.ExpEqToken:      js.OpAssign,
	js.AddEqToken:      js.OpAssign,
	js.SubEqToken:      js.OpAssign,
	js.LtLtEqToken:     js.OpAssign,
	js.GtGtEqToken:     js.OpAssign,
	js.GtGtGtEqToken:   js.OpAssign,
	js.BitAndEqToken:   js.OpAssign,
	js.BitXorEqToken:   js.OpAssign,
	js.BitOrEqToken:    js.OpAssign,
	js.ExpToken:        js.OpExp,
	js.MulToken:        js.OpMul,
	js.DivToken:        js.OpMul,
	js.ModToken:        js.OpMul,
	js.AddToken:        js.OpAdd,
	js.SubToken:        js.OpAdd,
	js.LtLtToken:       js.OpShift,
	js.GtGtToken:       js.OpShift,
	js.GtGtGtToken:     js.OpShift,
	js.LtToken:         js.OpCompare,
	js.LtEqToken:       js.OpCompare,
	js.GtToken:         js.OpCompare,
	js.GtEqToken:       js.OpCompare,
	js.InToken:         js.OpCompare,
	js.InstanceofToken: js.OpCompare,
	js.EqEqToken:       js.OpEquals,
	js.NotEqToken:      js.OpEquals,
	js.EqEqEqToken:     js.OpEquals,
	js.NotEqEqToken:    js.OpEquals,
	js.BitAndToken:     js.OpBitAnd,
	js.BitXorToken:     js.OpBitXor,
	js.BitOrToken:      js.OpBitOr,
	js.AndToken:        js.OpAnd,
	js.OrToken:         js.OpOr,
	js.NullishToken:    js.OpCoalesce,
	js.CommaToken:      js.OpExpr,
}

func exprPrec(i js.IExpr) js.OpPrec { _ = "STUB: not implemented"; return *new(js.OpPrec) }

// CommaExpr

func hasSideEffects(i js.IExpr) bool {
	_ = "STUB: not implemented"
	// assume that variable usage and that the index operator themselves have no side effects
	return false
}

// TODO: use in more cases
func groupExpr(i js.IExpr, prec js.OpPrec) js.IExpr {
	_ = "STUB: not implemented"
	return *new(js.IExpr)
}

// TODO: use in more cases
func condExpr(cond, x, y js.IExpr) js.IExpr { _ = "STUB: not implemented"; return *new(js.IExpr) }

func commaExpr(x, y js.IExpr) js.IExpr { _ = "STUB: not implemented"; return *new(js.IExpr) }

func innerExpr(i js.IExpr) js.IExpr { _ = "STUB: not implemented"; return *new(js.IExpr) }

func finalExpr(i js.IExpr) js.IExpr { _ = "STUB: not implemented"; return *new(js.IExpr) }

// return first

func isTrue(i js.IExpr) bool { _ = "STUB: not implemented"; return false }

func isFalse(i js.IExpr) bool { _ = "STUB: not implemented"; return false }

func isEqualExpr(a, b js.IExpr) bool { _ = "STUB: not implemented"; return false }

// TODO: use reflect.DeepEqual?

func toNullishExpr(condExpr *js.CondExpr) (js.IExpr, bool) {
	_ = "STUB: not implemented"
	return *new(js.IExpr), false
}

// convert conditional expression to nullish:  a==null?b:a  =>  a??b

// convert conditional expression to optional expr:  a==null?undefined:a.b  =>  a?.b

func isUndefinedOrNullVar(i js.IExpr) (*js.Var, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func isUndefinedOrNull(i js.IExpr) bool { _ = "STUB: not implemented"; return false }

func isUndefined(i js.IExpr) bool { _ = "STUB: not implemented"; return false }

// TODO: only if not defined

// returns whether truthy and whether it could be coerced to a boolean (i.e. when returns (false,true) this means it is falsy)
func isTruthy(i js.IExpr) (bool, bool) { _ = "STUB: not implemented"; return false, false }

// returns whether falsy and whether it could be coerced to a boolean (i.e. when returns (false,true) this means it is truthy)
func isFalsy(i js.IExpr) (bool, bool) { _ = "STUB: not implemented"; return false, false }

// falsy

// truthy

// truthy

// falsy

// falsy

// falsy

// unknown

func isBooleanExpr(expr js.IExpr) bool { _ = "STUB: not implemented"; return false }

func invertBooleanOp(op js.TokenType) js.TokenType {
	_ = "STUB: not implemented"
	return *new(js.TokenType)
}

func optimizeBooleanExpr(expr js.IExpr, invert bool, prec js.OpPrec) js.IExpr {
	_ = "STUB: not implemented"

	// unary !(boolean) has already been handled
	return *new(js.IExpr)
}

func optimizeUnaryExpr(expr *js.UnaryExpr, prec js.OpPrec) js.IExpr {
	_ = "STUB: not implemented"
	return *new(js.IExpr)
}

// rewrite !(a||b) to !a&&!b
// rewrite !(a==0||b==0) to a!=0&&b!=0
// savings if rewritten (group parentheses and not-token)

// add two not-tokens for left and right

// == and === can become != and !==

// add group if it wasn't already there

// remove group

func (m *jsMinifier) optimizeCondExpr(expr *js.CondExpr, prec js.OpPrec) js.IExpr {
	_ = "STUB: not implemented"
	// remove double negative !! in condition, or switch cases for single negative !
	return *new(js.IExpr)
}

// if condition is truthy

// if condition is falsy

// if condition is equal to true body
// for higher prec we need to add group parenthesis, and for lower prec we have parenthesis anyways. This only is shorter if len(expr.X) >= 3. isEqualExpr only checks for literal variables, which is a name will be minified to a one or two character name.

// if condition is equal to false body
// for higher prec we need to add group parenthesis, and for lower prec we have parenthesis anyways. This only is shorter if len(expr.X) >= 3. isEqualExpr only checks for literal variables, which is a name will be minified to a one or two character name.

// if true and false bodies are equal

// no need to check whether left/right need to add groups, as the space saving is always more

// recompress the conditional expression inside

// shorten when true and false bodies are true and false

// trueX != trueY

// falseX != falseY

// nested conditional expression with same false bodies

// regular conditional expression
// convert  (a,b)?c:d  =>  a,b?c:d

// recompress the conditional expression inside

func isHexDigit(b byte) bool { _ = "STUB: not implemented"; return false }

func mergeBinaryExpr(expr *js.BinaryExpr) {
	_ = "STUB: not implemented"
	// merge string concatenations which may be intertwined with other additions
	return
}

// limit recursion

// unescaped quotes will be repaired in minifyString later on

func minifyString(b []byte, allowTemplate bool) []byte { _ = "STUB: not implemented"; return nil }

// switch quotes if more optimal

// default to " for better GZIP compression

// strip unnecessary escapes

func replaceEscapes(b []byte, quote byte, prefix, suffix int) []byte {
	_ = "STUB: not implemented"
	// strip unnecessary escapes
	return nil
}

// keep escape sequence

// number of characters to skip

// line continuations

// don't convert \x00 to \0 if it may be an octal number
// hexadecimal escapes

// don't convert NULL to literal NULL (gives JS parsing problems)

// don't convert NULL to \0 (may be an octal number)

// decode unicode character to UTF-8 and put at the end of the escape sequence
// then skip the first part of the escape sequence until the decoded character

// octal escapes (legacy), \0 already handled (quote != `)

// remove unnecessary escape character, anything but 0x00, 0x0A, 0x0D, \, ' or "

// may not be escaped properly when changing quotes

// avoid append

// was overwritten above

// avoid append

// was overwritten above

var regexpEscapeTable = [256]bool{
	// ASCII
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, true, false, false, false, // $
	true, true, true, true, false, false, true, true, // (, ), *, +, ., /
	true, true, true, true, true, true, true, true, // 0, 1, 2, 3, 4, 5, 6, 7
	true, true, false, false, false, false, false, true, // 8, 9, ?

	false, false, true, false, true, false, false, false, // B, D
	false, false, false, false, false, false, false, false,
	true, false, false, true, false, false, false, true, // P, S, W
	false, false, false, true, true, true, true, false, // [, \, ], ^

	false, false, true, true, true, false, true, false, // b, c, d, f
	false, false, false, true, false, false, true, false, // k, n
	true, false, true, true, true, true, true, true, // p, r, s, t, u, v, w
	true, false, false, true, true, true, false, false, // x, {, |, }

	// non-ASCII
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
}

var regexpClassEscapeTable = [256]bool{
	// ASCII
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	true, true, true, true, true, true, true, true, // 0, 1, 2, 3, 4, 5, 6, 7
	true, true, false, false, false, false, false, false, // 8, 9

	false, false, false, false, true, false, false, false, // D
	false, false, false, false, false, false, false, false,
	true, false, false, true, false, false, false, true, // P, S, W
	false, false, false, false, true, true, false, false, // \, ]

	false, false, true, true, true, false, true, false, // b, c, d, f
	false, false, false, false, false, false, true, false, // n
	true, false, true, true, true, true, true, true, // p, r, s, t, u, v, w
	true, false, false, false, false, false, false, false, // x

	// non-ASCII
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,

	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false,
}

func minifyRegExp(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func removeUnderscoresAndSuffix(b []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func decimalNumber(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }

func binaryNumber(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }

func octalNumber(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }

func hexadecimalNumber(num []byte, prec int) []byte { _ = "STUB: not implemented"; return nil }
