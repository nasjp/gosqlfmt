package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	input     = os.Stdin
	outputOK  = os.Stdout
	outputErr = os.Stderr
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(outputErr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	mark, files, write, err := ParseFlags()
	if err != nil {
		return err
	}

	for _, file := range files {
		src, err := ioutil.ReadFile(file)
		if err != nil {
			return newErr(err, "read file failed")
		}

		out, err := FormatSQL(src, file, mark)
		if err != nil {
			return err
		}

		if err := Output(out, file, write); err != nil {
			return err
		}
	}

	return nil
}
