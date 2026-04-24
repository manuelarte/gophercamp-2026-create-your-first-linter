package unexportedconstantscheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "unexportedconstantscheck",
		Doc:      "unexportedconstantscheck checks if unexported constants starts with _",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, isGenDecl := decl.(*ast.GenDecl)
			if !isGenDecl {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, isValueSpec := spec.(*ast.ValueSpec)
				if !isValueSpec {
					continue
				}
				// TODO(manuelarte): to be done
				//nolint:forbidigo // remove later
				fmt.Printf("%+v\n", valueSpec)
				// End TODO(manuelarte)
			}
		}
	}

	//nolint:nilnil //any, error
	return nil, nil
}

//nolint:unused // to be used later
func newUnexportedConstantsCheckDiag(i *ast.Ident) analysis.Diagnostic {
	msg := fmt.Sprintf("unexported constant %q should be prefixed with _",
		i.Name)

	return analysis.Diagnostic{
		Pos:     i.Pos(),
		End:     i.End(),
		Message: msg,
		/* TODO(manuelarte): add later
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: msg,
				TextEdits: []analysis.TextEdit{
					{
						Pos: i.Pos(),
						End: i.End(),
						// TODO(manuelarte): Can someone see the problem of fixing it?
						NewText: fmt.Appendf(nil, "_%s", i.Name),
					},
				},
			},
		},
		*/
	}
}

//nolint:unused // to be used later
func newPrefixedErrConsCheckDiag(i *ast.Ident) analysis.Diagnostic {
	msg := fmt.Sprintf("unexported err constant %q should not be prefixed with _",
		i.Name)

	return analysis.Diagnostic{
		Pos:     i.Pos(),
		End:     i.End(),
		Message: msg,
		/* TODO(manuelarte): add later
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: msg,
				TextEdits: []analysis.TextEdit{
					{
						Pos: i.Pos(),
						End: i.End(),
						NewText: fmt.Appendf(nil, "%s", i.Name[1:]),
					},
				},
			},
		},
		*/
	}
}
