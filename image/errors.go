// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package image

import "fmt"

var (
	ErrFormatNotSupported = fmt.Errorf("format not supported")
	ErrBytesLimitExceeded = fmt.Errorf("available bytes limit exceeded")
	ErrOpeningFile        = fmt.Errorf("error opening file")
	ErrDecodingImage      = fmt.Errorf("error decoding image")
	ErrWrittingFile       = fmt.Errorf("error writing file")
)

// wrap function helps to create nested errors supporting string formating and
// error comparation using errors.Is().
func wrap(parent error, tmp string, args ...interface{}) error {
	return fmt.Errorf("%w: %w", parent, fmt.Errorf(tmp, args...))
}
