package flscanner

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Scanner struct {
	file  string
	inBuf *bytes.Buffer

	ch    rune
	eof   bool
	token *Token

	offset int

	Line   int
	Column int

	Errors []error
}

func NewScanner(path string) *Scanner {
	var s Scanner

	s.file = path
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		s.raise(err)
		return &s
	}

	buf = append(buf, '\n') // just to be sure

	s.inBuf = bytes.NewBuffer(buf)

	s.eof = false
	s.Line = 1
	s.Column = 1
	s.offset = 0

	return &s
}

const (
	eofRune = -1
	errRune = -2
)

func (s *Scanner) peek() rune {
	if s.eof {
		return eofRune
	}

	r, _, err := s.inBuf.ReadRune()
	defer s.inBuf.UnreadRune() // rewind

	if err == io.EOF {
		return eofRune
	} else if err != nil {
		s.raise(err)
		return errRune
	}

	return r
}

func (s *Scanner) next() rune {
	if s.eof {
		return eofRune
	}

	r, n, err := s.inBuf.ReadRune()
	if err == io.EOF {
		s.eof = true
		s.ch = eofRune
		s.offset += n
		s.Column++
		return eofRune
	} else if err != nil {
		s.raise(err)
		return errRune
	}

	if r == '\n' {
		s.Column = 0
		s.Line++
	}

	s.offset += n
	s.Column++
	s.ch = r

	return r
}

func (s *Scanner) match(r rune) bool {
	if s.peek() == r {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) matchIf(f func(rune) bool) bool {
	if f(s.peek()) {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) matchIfString(f func(string) bool) bool {
	if f(s.token.Text + string(s.peek())) {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) Next() (token *Token) {
	s.token = &Token{
		Line:   s.Line,
		Column: s.Column,
	}
	token = s.token

	if s.match(eofRune) {
		token.Kind = EOF
		s.next()
		return
	}

	switch {
	case s.matchIf(isSpace):
	case s.match('('):
		if s.match('*') {
			s.scanBlockComment()
			return s.Next()
		}
		fallthrough // '('
	case s.matchIf(isDelimiter):
		token.Text = string(s.ch)
		if s.matchIfString(isOperator) {
			token.Text += string(s.ch)
			token.Kind = reserved[token.Text]
			return
		}
		token.Kind = TokenKind(s.ch)
		return
	case s.match('"'):
		s.scanString()
		return
	case s.matchIf(isNameStart):
		s.scanName()
		return
	case s.matchIf(isDigit):
		s.scanNumber()
		return
	case s.matchIf(isUtf8):
		s.errorf("unexpected UTF-8 character")
		token.Text += string(s.ch)
		token.Kind = BAD
		s.next()
		return
	default:
		s.errorf("malformed UTF-8 encoding")
		s.next()
	}

	return s.Next()
}

func (s *Scanner) scanNumber() {
	s.token.Text = string(s.ch)
	s.token.Kind = INT

	if s.ch == '0' && s.match('x') {
		s.token.Text = "0x"
		for s.matchIf(isHexDigit) {
			s.token.Text += string(s.ch)
		}

		if len(s.token.Text) == 2 {
			s.errorf("malformed hexadecimal literal")
			s.token.Kind = BAD
			return
		}

		return
	}

	for s.matchIf(isDigit) {
		s.token.Text += string(s.ch)
	}

	if s.match('.') {
		s.token.Text += "."
		s.token.Kind = FLOAT
		for s.matchIf(isDigit) {
			s.token.Text += string(s.ch)
		}
		return
	}
}

func (s *Scanner) scanString() {
	s.token.Kind = STRING
	s.token.Text = ""
	for !s.eof {
		switch {
		case s.eof || s.match('\n'):
			s.errorf("unclosed string literal")
			return
		case s.matchIf(unicode.IsControl):
			s.errorf("illegal control character in string literal")
		case s.match('"'):
			return
		case s.match('\\'):
			s.scanEscape()
		default:
			s.next()
			s.token.Text += string(s.ch)
		}
	}
}

func (s *Scanner) scanEscape() bool {
	// escape slash is already matched
	switch s.next() {
	case 'n':
		s.token.Text += "\n"
	case 'r':
		s.token.Text += "\r"
	case 't':
		s.token.Text += "\t"
	case '\\':
		s.token.Text += "\\"
	case '\'':
		s.token.Text += "'"
	case '"':
		s.token.Text += "\""
	case 'u': // unicode
		if !s.match('{') {
			s.errorf("missing opening '{' in unicode escape sequence")
			s.token.Text += "\\u"
			return false
		}

		esc := ""
		for s.matchIf(isHexDigit) {
			esc += string(s.ch)
		}

		switch {
		case len(esc) == 0 && s.match('}'):
			s.errorf("empty unicode escape sequence")
			s.token.Text += "\\u{}"
			return false
		case len(esc) == 0 && !s.match('}'):
			rtext := safeRune(s.peek())
			s.errorf("unexpected character in unicode escape sequence '%s'", rtext)
			s.token.Text += "\\u" + rtext
			return false
		case len(esc) > 0 && !s.match('}'):
			s.errorf("missing closing '}' in unicode escape sequence")
			s.token.Text += "\\u{" + esc
			return false
		}

		n, err := strconv.ParseInt(esc, 16, 0)
		if err != nil {
			s.raise(err)
			s.token.Text += "\\u{" + esc + "}"
			return false
		}

		u := make([]byte, 4)
		utf8.EncodeRune(u, rune(n))
		s.token.Text += string(bytes.Trim(u, "\x00")) // remove NULL bytes
	default: // hexadecimal
		if !isHexDigit(s.ch) {
			rtext := safeRune(s.ch)
			s.errorf("unexpected character in hexadecimal escape sequence '%s'", rtext)
			s.token.Text += "\\" + rtext
			return false
		}

		esc := string(s.ch)
		if !s.matchIf(isHexDigit) {
			rtext := safeRune(s.peek())
			s.errorf("unexpected character in hexadecimal escape sequence '%s'", rtext)
			s.token.Text += "\\" + esc + rtext
			return false
		}
		esc += string(s.ch)

		n, err := strconv.ParseInt(esc, 16, 0)
		if err != nil {
			s.raise(err)
			s.token.Text += "\\" + esc
			return false
		}
		s.token.Text += string(n)
	}
	return true
}

func (s *Scanner) scanName() {
	s.token.Text += string(s.ch)
	for s.matchIf(isNamePart) {
		s.token.Text += string(s.ch)
	}

	if k, ok := reserved[s.token.Text]; ok {
		s.token.Kind = k
	} else {
		s.token.Kind = NAME
	}
}

func (s *Scanner) scanBlockComment() {
	for depth := 1; depth > 0; {
		switch {
		case s.eof:
			s.errorf("unclosed block comment")
			return
		case s.match('(') && s.match('*'):
			depth++
		case s.match('*') && s.match(')'):
			depth--
		default:
			s.next()
		}
	}
}

const scanErrPrefix = "error: "

// errorf generates a new scanner error appended to the scanner's Errors field
func (s *Scanner) errorf(fmtStr string, args ...interface{}) {
	pfx := fmt.Sprintf("%s ~ line %d, column %d\n  => ", s.file, s.token.Line, s.token.Column)
	err := fmt.Errorf(scanErrPrefix+pfx+fmtStr+"\n", args...)
	s.Errors = append(s.Errors, err)
}

// raise directly promote any error to a printable Scanner error
func (s *Scanner) raise(err error) {
	err2 := fmt.Errorf(scanErrPrefix+"%s\n  %s\n", s.file, err.Error())
	s.Errors = append(s.Errors, err2)
}

func isUtf8(r rune) bool {
	return utf8.ValidRune(r)
}

func isSpace(r rune) bool {
	switch r {
	case ' ', '\n', '\r', '\t':
		return true
	default:
		return false
	}
}

func isOperator(s string) bool {
	switch s {
	case "!", "~", "+", "-", "*", "/", "%", "&", "|", "^", "=", "<", ">":
		return true
	case "++", "::", "&&", "||", "==", "<=", ">=", "!=", "->", "=>":
		return true
	case ">>", "<<", ":=":
		return true
	}
	return false
}

func isDelimiter(r rune) bool {
	return !isSpace(r) && strings.IndexRune("<>()[]{},;:._+-*/%^&|=#!?", r) >= 0
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isHexDigit(r rune) bool {
	if isDigit(r) {
		return true
	}
	return (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}

func isNameStart(r rune) bool {
	return !isSpace(r) && !isDelimiter(r) && !isDigit(r)
}

func isNamePart(r rune) bool {
	return !isSpace(r) && strings.IndexRune("\"();", r) < 0
}

func safeRune(r rune) string {
	switch r {
	case '\n':
		return "\\n"
	case '\r':
		return "\\r"
	case '\t':
		return "\\t"
	case '\'':
		return "\\'"
	case '"':
		return "\\\""
	default:
		if unicode.IsPrint(r) {
			return string(r)
		}
		return fmt.Sprintf("\\%x", r)
	}
}
