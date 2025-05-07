// Package errorsis is the library used to implement the errorsis command-line tool.
package errorsis

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var NoErrorIsStruct = &analysis.Analyzer{
	Name: "noerrorsisstruct",
	Doc:  "Avoid using errors.Is() with struct or struct pointer",
	Run:  noErrorIsStruct,
}

//nolint:gocognit,ireturn
func noErrorIsStruct(pass *analysis.Pass) (interface{}, error) {
	var errorsIsType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() == "errors" {
			errorsIsType = pkg.Scope().Lookup("Is").Type()

			break
		}
	}

	if errorsIsType == nil {
		return nil, nil //nolint:nilnil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			var valid bool

			var call *ast.CallExpr

			call, valid = node.(*ast.CallExpr)
			if !valid || len(call.Args) != 2 {
				return true
			}

			if pass.TypesInfo.TypeOf(call.Fun) != errorsIsType {
				return true
			}

			arg := call.Args[1]

			var unaryExpr *ast.UnaryExpr

			// Only flag composite literals like &SomeStruct{}
			unaryExpr, valid = arg.(*ast.UnaryExpr)
			if !valid || unaryExpr.Op != token.AND {
				return true
			}

			_, valid = unaryExpr.X.(*ast.CompositeLit)
			if !valid {
				return true
			}

			argType := pass.TypesInfo.TypeOf(arg)
			errorIface := types.Universe.Lookup("error").Type().Underlying().(*types.Interface) //nolint:forcetypeassert

			if !types.Implements(argType, errorIface) {
				return true
			}

			if ptrType, ok := argType.(*types.Pointer); ok {
				elem := ptrType.Elem()

				named, ok := elem.(*types.Named)
				if !ok {
					return true
				}

				if !types.Implements(named, errorIface) {
					pass.Reportf(arg.Pos(), "incorrect usage of errors.Is: &T{} used where only *T implements error")
				}
			}

			return true
		})
	}

	return nil, nil //nolint:nilnil
}
