package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func FormatSQL(src []byte, fileName string, mark string) ([]byte, error) {

	fset := token.NewFileSet()
	parserMode := parser.ParseComments

	astFile, err := parser.ParseFile(fset, fileName, src, parserMode)
	if err != nil {
		return nil, newErr(err, "parser.ParseFile failed")
	}

	fmtSQL(astFile, fset, mark)

	buf := bytes.NewBuffer(nil)

	if err := format.Node(buf, fset, astFile); err != nil {
		return nil, newErr(err, "format.Node failed")
	}

	return buf.Bytes(), nil
}

func fmtSQL(f *ast.File, fset *token.FileSet, mark string) {
	commentMap := ast.NewCommentMap(fset, f, f.Comments)
	for n, commentGroups := range commentMap {
		for _, commentGroup := range commentGroups {
			for _, comment := range commentGroup.List {
				if comment.Text == mark {
					queryInspect(n, fset, fset.File(comment.Slash).Line(comment.Slash))
				}
			}
		}
	}
}

func queryInspect(f ast.Node, fset *token.FileSet, commentLine int) {
	ast.Inspect(f, func(n ast.Node) bool {
		if basicLit, ok := n.(*ast.BasicLit); ok {
			if basicLit.Kind == token.STRING && len(basicLit.Value) > 2 {
				litLine := fset.File(basicLit.End()).Line(basicLit.End())
				if litLine == commentLine {
					// TODO implement format sql
					// basicLit.Value = ParseSQL(basicLit.Value).Tokenize().Format().String()
				}
			}
		}

		return true
	})
}
