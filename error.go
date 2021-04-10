package main

import "github.com/pkg/errors"

func newErr(err error, message string) error {
	return errors.Wrap(err, message)
}
