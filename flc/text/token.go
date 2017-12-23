package flscanner

import "fmt"

type Token struct {
	Text   string
	Kind   TokenKind
	Line   int
	Column int
}

type TokenKind int

const (
	BAD TokenKind = iota
	EOF

	VALUE_TYPE TokenKind = iota + 0xff
	NAME

	UNIT
	BOOL
	INT
	FLOAT
	STRING

	TYPE
	FUNC
	LET
	VAR
	IF
	THEN
	ELSE
	IN
	WHERE

	SHL
	SHR
	DECL
	CONCAT
	BLOCK
	ARROW

	EQ
	NE
	AND
	OR
	LE
	GE
)

var reserved = map[string]TokenKind{
	"unit":   VALUE_TYPE,
	"bool":   VALUE_TYPE,
	"int":    VALUE_TYPE,
	"float":  VALUE_TYPE,
	"string": VALUE_TYPE,

	"true":  BOOL,
	"false": BOOL,

	"type":  TYPE,
	"func":  FUNC,
	"let":   LET,
	"var":   VAR,
	"if":    IF,
	"then":  THEN,
	"else":  ELSE,
	"in":    IN,
	"where": WHERE,

	"<<": SHL,
	">>": SHR,
	":=": DECL,
	"++": CONCAT,
	"::": BLOCK,
	"->": ARROW,
	"=>": ARROW,
	"==": EQ,
	"!=": NE,
	"&&": AND,
	"||": OR,
	"<=": LE,
	">=": GE,
}

var tokenName = [...]string{
	BAD:        "BAD",
	EOF:        "EOF",
	VALUE_TYPE: "VALUE_TYPE",
	NAME:       "NAME",
	UNIT:       "UNIT",
	BOOL:       "BOOL",
	INT:        "INT",
	FLOAT:      "FLOAT",
	STRING:     "STRING",
	TYPE:       "TYPE",
	FUNC:       "FUNC",
	LET:        "LET",
	VAR:        "VAR",
	IF:         "IF",
	THEN:       "THEN",
	ELSE:       "ELSE",
	IN:         "IN",
	WHERE:      "WHERE",
	SHL:        "SHL",
	SHR:        "SHR",
	DECL:       "DECL",
	CONCAT:     "CONCAT",
	BLOCK:      "BLOCK",
	ARROW:      "ARROW",
	EQ:         "EQ",
	NE:         "NE",
	AND:        "AND",
	OR:         "OR",
	LE:         "LE",
	GE:         "GE",
}

func (t *Token) String() string {
	// Don't print unwanted characters (eg. newline)
	safe := ""
	for _, r := range t.Text {
		safe += safeRune(r)
	}

	switch {
	case t.Kind == EOF:
		return "EOF"
	case t.Kind == STRING:
		return fmt.Sprintf("STRING(\"%s\")", safe)
	case t.Kind == BAD:
		return fmt.Sprintf("BAD('%s')", safe)
	case t.Kind < VALUE_TYPE:
		return fmt.Sprintf("'%s'", safe)
	default:
		return fmt.Sprintf("%s('%s')", tokenName[t.Kind], safe)
	}
}
