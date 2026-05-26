package js

import (
	"github.com/tdewolff/parse/v2/js"
)

func optimizeStmt(i js.IStmt) js.IStmt {
	_ = "STUB: not implemented"
	// convert if/else into expression statement, and optimize blocks
	return *new(js.IStmt)
}

// TODO: remove if and return StmtList(Cond, Body)

// falsy

// TODO: remove if and return StmtList(Cond, Body)

// TODO: remove function name in var name=function name(){}
//for _, item := range decl.List {
//	if v, ok := item.Binding.(*js.Var); ok && item.Default != nil {
//		if fun, ok := item.Default.(*js.FuncDecl); ok && fun.Name != nil && bytes.Equal(v.Data, fun.Name.Data) {
//			scope := fun.Body.Scope
//			for i, vorig := range scope.Declared {
//				if fun.Name == vorig {
//					scope.Declared = append(scope.Declared[:i], scope.Declared[i+1:]...)
//				}
//			}
//			scope.AddUndeclared(v)
//			v.Uses += fun.Name.Uses - 1
//			fun.Name.Link = v
//			fun.Name = nil
//		}
//	}
//}

// convert hoisted var declaration to expression or empty (if there are no defines) statement

// TODO: remove unused declarations
//for i := 0; i < len(decl.List); i++ {
//	if v, ok := decl.List[i].Binding.(*js.Var); ok && v.Uses < 2 {
//		decl.List = append(decl.List[:i], decl.List[i+1:]...)
//		i--
//	}
//}
//if len(decl.List) == 0 {
//	return &js.EmptyStmt{}
//}

// merge body and remove braces if it is not a lexical declaration

// remove let or const declaration in otherwise empty scope, but keep assignments

func optimizeStmtList(list []js.IStmt, blockType blockType) []js.IStmt {
	_ = "STUB: not implemented"
	// merge expression statements as well as if/else statements followed by flow control statements
	return nil
}

// write index
// read index

// if(a)return b;else c  =>  if(a)b; c

// if body ends in flow statement (return, throw, break, continue), we can remove the else statement and put its body in the current scope

// merge expression statements with expression, return, and throw statements

// TODO: only merge lhs expression that don't have 'in' or 'of' keywords (slow to check?)

// this is the second VarDecl, so we are hoisting var declarations, which means the forInit variables are already in 'left'

// TODO: only merge lhs expression that don't have 'in' or 'of' keywords (slow to check?)

// merge const and let declarations, or non-hoisted var declarations

// remove from vardecls list of scope

// pull in assignments to variables into the declaration, e.g. var a;a=5  =>  var a=5

// TODO: only merge lhs expression that don't have 'in' or 'of' keywords (slow to check?)

// this is the second VarDecl, so we are hoisting var declarations, which means the forInit variables are already in 'left'

// TODO: only merge lhs expression that don't have 'in' or 'of' keywords (slow to check?)

// merge if/else with return/throw when followed by return/throw

// separate from expression merging in case of:  if(a)return b;b=c;return d

// either the if body is empty or the else body is empty. In case where both bodies have return/throw, we already rewrote that if statement to an return/throw statement

// remove superfluous return or continue

// rewrite function f(){return a,void 0} => function f(){a}
