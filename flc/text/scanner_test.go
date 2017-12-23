package flscanner

import (
	"flag"
	"fmt"
	"testing"
)

var file = flag.String("file", "", "the file to test")

func TestScanner(t *testing.T) {
	if *file == "" {
		t.Errorf("error: no input file")
	}

	s := NewScanner(*file)
	if len(s.Errors) > 0 {
		fmt.Println(s.Errors[0])
		return
	}

	var tok *Token
	tok = s.Next()
	for tok.Kind != EOF {
		fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok)
		tok = s.Next()
	}
	fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok.String())

	for _, err := range s.Errors {
		fmt.Print(err)
	}

	if len(s.Errors) > 0 {
		t.Errorf("scanner: failed with %d error(s)", len(s.Errors))
	}
}
