
// Package token is generated by GoGLL. Do not edit
package token

import(
	"fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
	typ        Type
	lext, rext int
	input      []rune
}

/*
New returns a new token.
lext is the left extent and rext the right extent of the token in the input.
input is the input slice scanned by the lexer.
*/
func New(t Type, lext, rext int, input []rune) *Token {
	return &Token{
		typ:   t,
		lext:  lext,
		rext:  rext,
		input: input,
	}
}

// GetLineColumn returns the line and column of the left extent of t
func (t *Token) GetLineColumn() (line, col int) {
	line, col = 1, 1
	for j := 0; j < t.lext; j++ {
		switch t.input[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetInput returns the input from which t was parsed.
func (t *Token) GetInput() []rune {
	return t.input
}

// Lext returns the left extent of t
func (t *Token) Lext() int {
	return t.lext
}

// Literal returs the literal runes of t scanned by the lexer
func (t *Token) Literal() []rune {
	return t.input[t.lext:t.rext]
}

// LiteralString returns string(t.Literal())
func (t *Token) LiteralString() string {
	return string(t.Literal())
}

// Rext returns the right extent of t in the input
func (t *Token) Rext() int {
	return t.rext
}

func (t *Token) String() string {
	return fmt.Sprintf("%s (%d,%d) %s",
		t.TypeID(), t.lext, t.rext, t.LiteralString())
}

// Type returns the token Type of t
func (t *Token) Type() Type {
	return t.typ
}

// TypeID returns the token Type ID of t. 
// This may be different from the literal of token t.
func (t *Token) TypeID() string {
	return t.Type().ID()
}

// Type is the token type
type Type int

func (t Type) String() string {
	return TypeToString[t]
}

// ID returns the token type ID of token Type t
func (t Type) ID() string {
	return TypeToID[t]
}


const(
	Error  Type = iota  // Error 
	EOF  // EOF 
	Type0  // ( 
	Type1  // ) 
	Type2  // . 
	Type3  // : 
	Type4  // ; 
	Type5  // < 
	Type6  // > 
	Type7  // [ 
	Type8  // ] 
	Type9  // any 
	Type10  // char_lit 
	Type11  // empty 
	Type12  // letter 
	Type13  // lowcase 
	Type14  // not 
	Type15  // nt 
	Type16  // number 
	Type17  // package 
	Type18  // string_lit 
	Type19  // tokid 
	Type20  // upcase 
	Type21  // { 
	Type22  // | 
	Type23  // } 
)

var TypeToString = []string{ 
	"Error",
	"EOF",
	"Type0",
	"Type1",
	"Type2",
	"Type3",
	"Type4",
	"Type5",
	"Type6",
	"Type7",
	"Type8",
	"Type9",
	"Type10",
	"Type11",
	"Type12",
	"Type13",
	"Type14",
	"Type15",
	"Type16",
	"Type17",
	"Type18",
	"Type19",
	"Type20",
	"Type21",
	"Type22",
	"Type23",
}

var StringToType = map[string] Type { 
	"Error" : Error, 
	"EOF" : EOF, 
	"Type0" : Type0, 
	"Type1" : Type1, 
	"Type2" : Type2, 
	"Type3" : Type3, 
	"Type4" : Type4, 
	"Type5" : Type5, 
	"Type6" : Type6, 
	"Type7" : Type7, 
	"Type8" : Type8, 
	"Type9" : Type9, 
	"Type10" : Type10, 
	"Type11" : Type11, 
	"Type12" : Type12, 
	"Type13" : Type13, 
	"Type14" : Type14, 
	"Type15" : Type15, 
	"Type16" : Type16, 
	"Type17" : Type17, 
	"Type18" : Type18, 
	"Type19" : Type19, 
	"Type20" : Type20, 
	"Type21" : Type21, 
	"Type22" : Type22, 
	"Type23" : Type23, 
}

var TypeToID = []string { 
	"Error", 
	"EOF", 
	"(", 
	")", 
	".", 
	":", 
	";", 
	"<", 
	">", 
	"[", 
	"]", 
	"any", 
	"char_lit", 
	"empty", 
	"letter", 
	"lowcase", 
	"not", 
	"nt", 
	"number", 
	"package", 
	"string_lit", 
	"tokid", 
	"upcase", 
	"{", 
	"|", 
	"}", 
}
