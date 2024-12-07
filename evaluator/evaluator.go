package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatement(node.Statements)
  case *ast.ExpressionStatement:
    return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatement(stmt []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmt {
		result = Eval(statement)
	}

	return result
}
