package main

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	sqlStmtMark = "/*SQL*/"
	skipDir     = "testdata"
)

func ParseFlags() (sqlStmtCommentMark string, files []string, write bool, err error) {
	commentFlag := flag.String("comment", sqlStmtMark, "set sql statement mark if ")
	writeFlag := flag.Bool("w", false, "overwrite go file")

	flag.Func("target", "go files or dir with comma separator", func(s string) error {
		files = make([]string, 0)
		for _, target := range strings.Split(s, ",") {
			fs, err := parseGoFiles(target)
			if err != nil {
				return err
			}

			files = append(files, fs...)
		}

		return nil
	})

	flag.Parse()

	if files == nil {
		files, err = parseGoFiles(".")
		if err != nil {
			return "", nil, false, err
		}
	}

	return *commentFlag, files, *writeFlag, nil
}

func parseGoFiles(target string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(filepath.Clean(target), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// for _, s := range strings.Split(path, "/") {
		// 	if s == skipDir {
		// 		return nil
		// 	}
		// }

		if isGoFile(info) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, newErr(err, "parse file failed")
	}

	return files, nil
}

func isGoFile(info fs.FileInfo) bool {
	name := info.Name()
	return !info.IsDir() && len(name) > 0 && name[0] != '.' && filepath.Ext(name) == ".go"
}
