package main

import (
	"fmt"
	"os"
)

func Output(src []byte, fileName string, write bool) error {
	if !write {
		fmt.Fprintf(outputOK, "%s", src)

		return nil
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return newErr(err, "open file failed")
	}

	defer f.Close()

	if _, err := f.Write(src); err != nil {
		return err
	}

	return nil
}
