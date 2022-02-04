package parser

import (
	"io"
)

/** */
type Node struct {
}

type Parser struct {
	Scan *Scanner
}

/** Read tokens from the given scanner and produce parse tree 
If parsing fails, terminate exit code 1 **/
func (p *Parser) Parse(fout io.Writer, name string, debug bool) *Node {
	return nil
}
