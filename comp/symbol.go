package virt

import "fmt"

// A SymbolTable holds every symbols of a given scope
type SymbolTable struct {
	symbols map[string]Symbol
}

// A Symbol holds the type and the text of a symbol (identifier, label, ...)
type Symbol interface {
	// The symbol value
	Text() string

	fmt.Formatter

	// The symbol's type
	Type() Type

	// Whether the symbol has a value or not
	Initialized() bool

	// The symbol's value
	Value() fmt.Stringer

	fmt.Stringer
}
