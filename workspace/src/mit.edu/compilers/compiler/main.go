package main

import (
	"os"
	"path/filepath"
	"strings"
	"mit.edu/compilers/compiler/parser"
)

func main() {
	var err error

	cli := NewCLI()
	cli.ParseArgs(os.Args[1:], []string{})

	fout := os.Stdout
	if cli.outfile != "" {
		fout, err = os.Create(cli.outfile)
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	}

	infileShortName := "stdin"
	fin_ := os.Stdin
	if cli.infile != "" {
		fin_, err = os.Open(cli.infile)
		defer fin_.Close()
		if err != nil {
			panic(err)
		}
		infileShortName = strings.Split(filepath.Base(cli.infile), ".")[0]
	}
	
	scanner := &parser.Scanner{Fin: fin_}
	scanner.Scan(fout, infileShortName)
	if cli.target == CLI_TARGET_SCAN {
		return
	}
	p := &parser.Parser{Scan: scanner}
	p.Parse(fout, infileShortName, cli.debug)
}
