// Package js minifies ECMAScript 2021 following the language specification at https://tc39.es/ecma262/.
package js

import (
	"io"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/parse/v2/js"
)

type blockType int

const (
	defaultBlock blockType = iota
	functionBlock
	iterationBlock
)

// Minifier is a JS minifier.
type Minifier struct {
	Precision           int // number of significant digits
	KeepVarNames        bool
	useAlphabetVarNames bool
	Version             int
}

func (o *Minifier) minVersion(version int) bool { _ = "STUB: not implemented"; return false }

// Minify minifies JS data, it reads from r and writes to w.
func Minify(m *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Minify minifies JS data, it reads from r and writes to w.
func (o *Minifier) Minify(_ *minify.M, w io.Writer, r io.Reader, params map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

type expectExpr int

const (
	expectAny      expectExpr = iota
	expectExprStmt            // in statement
	expectExprBody            // in arrow function body
)

type jsMinifier struct {
	o *Minifier
	w io.Writer

	prev           []byte
	needsSemicolon bool       // write a semicolon if required
	needsSpace     bool       // write a space if next token is an identifier
	expectExpr     expectExpr // avoid ambiguous syntax such as an expression starting with function
	groupedStmt    bool       // avoid ambiguous syntax by grouping the expression statement
	inFor          bool
	spaceBefore    byte

	renamer *renamer
}

func (m *jsMinifier) write(b []byte) {
	_ = "STUB: not implemented"
	// 0 < len(b)
	return
}

func (m *jsMinifier) writeSpaceAfterIdent() {
	_ = "STUB: not implemented"
	// space after identifier and after regular expression (to prevent confusion with its tag)
	return
}

func (m *jsMinifier) writeSpaceBeforeIdent() { _ = "STUB: not implemented"; return }

func (m *jsMinifier) writeSpaceBefore(c byte) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) requireSemicolon() { _ = "STUB: not implemented"; return }

func (m *jsMinifier) writeSemicolon() { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyStmt(i js.IStmt) { _ = "STUB: not implemented"; return }

// prevent: if(a){if(b)c}else d;  =>  if(a)if(b)c;else d;

// can only be variable, function, or class decl

// bang comment

func (m *jsMinifier) minifyBlockStmt(stmt *js.BlockStmt) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyBlockAsStmt(blockStmt *js.BlockStmt) {
	_ = "STUB: not implemented"
	// minify block when statement is expected, i.e. semicolon if empty or remove braces for single statement
	// assume we already renamed the scope
	return
}

func (m *jsMinifier) minifyStmtOrBlock(i js.IStmt, blockType blockType) {
	_ = "STUB: not implemented"
	// minify stmt or a block
	return
}

// optimizeStmtList can in some cases expand one stmt to two shorter stmts

func (m *jsMinifier) minifyAlias(alias js.Alias) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyParams(params js.Params, removeUnused bool) {
	_ = "STUB: not implemented"
	// remove unused parameters from the end
	return
}

func (m *jsMinifier) minifyArguments(args js.Args) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyVarDecl(decl *js.VarDecl, onlyDefines bool) {
	_ = "STUB: not implemented"
	return
}

// remove 'var' when hoisting variables

func (m *jsMinifier) minifyFuncDecl(decl *js.FuncDecl, inExpr bool) {
	_ = "STUB: not implemented"
	// TODO: rewrite to arrow function if doe snot refer to this?
	//if !decl.Generator && decl.Name != nil && (!inExpr || 1 < decl.Name.Uses) {
	//	m.write(decl.Name.Data)
	//	m.write(equalBytes)
	//	m.minifyArrowFunc(&js.ArrowFunc{
	//		Async:  decl.Async,
	//		Params: decl.Params,
	//		Body:   decl.Body,
	//	})
	//	return
	//}
	return
}

// TODO: remove function name, really necessary?
//if decl.Name != nil && decl.Name.Uses == 1 {
//	scope := decl.Body.Scope
//	for i, vorig := range scope.Declared {
//		if decl.Name == vorig {
//			scope.Declared = append(scope.Declared[:i], scope.Declared[i+1:]...)
//		}
//	}
//}

func (m *jsMinifier) minifyClassElementName(name js.ClassElementName) {
	_ = "STUB: not implemented"
	return
}

func (m *jsMinifier) minifyMethodDecl(decl *js.MethodDecl) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyArrowFunc(decl *js.ArrowFunc) { _ = "STUB: not implemented"; return }

// add space after async in: async a => ...

// merge expression statements to final return statement, remove function body braces

// remove empty return

func (m *jsMinifier) minifyClassDecl(decl *js.ClassDecl) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyPropertyName(name js.PropertyName) { _ = "STUB: not implemented"; return }

func (m *jsMinifier) minifyProperty(property js.Property) {
	_ = "STUB: not implemented"
	// property.Name is always set in ObjectLiteral
	return
}

// add 'old-name:' before BindingName as the latter will be renamed

func (m *jsMinifier) minifyBindingElement(element js.BindingElement) {
	_ = "STUB: not implemented"
	return
}

func (m *jsMinifier) minifyBinding(ibinding js.IBinding) { _ = "STUB: not implemented"; return }

// item.Key is always set

// add 'old-name:' before BindingName as the latter will be renamed

func (m *jsMinifier) minifyExpr(i js.IExpr, prec js.OpPrec) { _ = "STUB: not implemented"; return }

// </script>/ => < /script>/

// convert (a,b)&&c into a,b&&c but not a=(b,c)&&d into a=(b,c&&d)

// TODO: has effect on GZIP?
//if expr.Op == js.EqEqToken || expr.Op == js.NotEqToken || expr.Op == js.EqEqEqToken || expr.Op == js.NotEqEqToken {
//	// switch a==const for const==a, such as typeof a=="undefined" for "undefined"==typeof a (GZIP improvement)
//	if _, ok := expr.Y.(*js.LiteralExpr); ok {
//		expr.X, expr.Y = expr.Y, expr.X
//	}
//}

// change a===null||a===undefined to a==null

// typeof a === "string"  =>  typeof a == "string"

// "string" === typeof a  =>  "string" == typeof a

// TODO: use truthy instead of true?

// TODO: use truthy instead of true?

// a=a+1  =>  ++a

// a=a-1  =>  --a

// TODO: may break implicit "evaluation" of variables? see #863
//} else if right.Op == js.AddToken || right.Op == js.SubToken || right.Op == js.MulToken || right.Op == js.DivToken || right.Op == js.ModToken || right.Op == js.ExpToken || right.Op == js.LtLtToken || right.Op == js.GtGtToken || right.Op == js.GtGtGtToken || right.Op == js.BitAndToken || right.Op == js.BitOrToken || right.Op == js.BitXorToken {
//	// a=a+b  =>  a+=b
//	m.minifyExpr(left, js.OpLHS)
//	m.write(right.Op.Bytes())
//	m.write(equalBytes)
//	m.minifyExpr(y, js.OpAssign)
//	break

// 0 < len(m.prev) always

// +++  =>  + ++

// ---  =>  - --

// //  =>  / /

// +++  =>  + ++

// ---  =>  - --
// <!--  =>  <! --

// !""  =>  !0

// !"string"  =>  !1

// !/regexp/  =>  !1

// !123  =>  !1 (except for !0)

// should never happen

// 0 < len(m.prev) always

// prevent previous integer

// new a() => (new a), when inside a Member, Call or OptChain expression

// new a() => new a

// Number(x) => +x

//} else {
//	if js.OpUnary < prec {
//		m.write(openParenBytes)
//	}
//	m.write(plusBytes)
//	m.minifyExpr(&js.GroupExpr{expr.Args.List[0].Value}, js.OpUnary)
//	if js.OpUnary < prec {
//		m.write(closeParenBytes)
//	}

// Math.pow(a,b) => a**b

// Math.trunc(x) => x|0

// Math.abs(x) => x<0?-x:x

// Math.sqrt(x) => x**.5

// happens in for statement or when vars were hoisted

// only happens in object literal
