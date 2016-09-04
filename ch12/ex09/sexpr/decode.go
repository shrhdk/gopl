package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

// Token is a common interface of s-expression tokens.
type Token interface{}

// Symbol represetns struct key name, and constants.
type Symbol string

// String represetns string value.
type String string

// Int represents integer value.
type Int int

// StartList represents (
type StartList struct{}

// EndList represetns )
type EndList struct{}

// A Decoder reads and decodes s-expressions values from an input stream.
type Decoder struct {
	lex *lexer
}

// NewDecoder returns a new decoder that reads from r.
// The decoder introduces its own buffering and may read data from r beyond the s-expression values requested.
func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{lex}
}

// Token reads next token of s-expression from base reader.
func (dec *Decoder) Token() (Token, error) {
	dec.lex.next()
	switch dec.lex.token {
	case scanner.Ident:
		return Symbol(dec.lex.text()), nil
	case scanner.String:
		s, err := strconv.Unquote(dec.lex.text())
		if err != nil {
			return nil, err
		}
		return String(s), nil
	case scanner.Int:
		i, err := strconv.Atoi(dec.lex.text())
		if err != nil {
			return nil, err
		}
		return Int(i), nil
	case '(':
		return &StartList{}, nil
	case ')':
		return &EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("Unsupported token type: %v", dec.lex.token)
	}
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }
