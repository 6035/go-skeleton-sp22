package parser

import (
	"fmt"
	"io"
)

type Scanner struct {
	Fin io.Reader
}

/** Prints a representation of each token scanned from the input file `Fin`
If scanning fails, terminate exit code 1 **/
func (s *Scanner) Scan(fout io.Writer, name string) {
	
}

/** Returns the next token in the input file `Fin` */
func (s *Scanner) Next() (tok *Token, err error, eof bool) {
	return nil, nil, false
}

/** Stop execution and report error **/
func Error(message string) {
	panic(fmt.Errorf(message))
}
