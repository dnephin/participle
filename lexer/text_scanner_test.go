package lexer

import (
	"strings"
	"testing"
	"text/scanner"

	"github.com/gotestyourself/gotestyourself/assert"
)

func TestLexer(t *testing.T) {
	lexer := LexString("hello world")
	helloPos := Position{Offset: 0, Line: 1, Column: 1}
	worldPos := Position{Offset: 6, Line: 1, Column: 7}
	eofPos := Position{Offset: 11, Line: 1, Column: 12}
	assert.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, lexer.Peek(0))
	assert.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, lexer.Peek(0))
	assert.Equal(t, Token{Type: scanner.Ident, Value: "hello", Pos: helloPos}, lexer.Next())
	assert.Equal(t, Token{Type: scanner.Ident, Value: "world", Pos: worldPos}, lexer.Peek(0))
	assert.Equal(t, Token{Type: scanner.Ident, Value: "world", Pos: worldPos}, lexer.Next())
	assert.Equal(t, Token{Type: scanner.EOF, Value: "", Pos: eofPos}, lexer.Peek(0))
	assert.Equal(t, Token{Type: scanner.EOF, Value: "", Pos: eofPos}, lexer.Next())
}

func TestLexString(t *testing.T) {
	lexer := LexString(`"hello\nworld"`)
	assert.Equal(t, lexer.Next(), Token{Type: scanner.String, Value: "hello\nworld", Pos: Position{Line: 1, Column: 1}})
}

func TestLexSingleString(t *testing.T) {
	lexer := LexString(`'hello\nworld'`)
	assert.Equal(t, lexer.Next(), Token{Type: scanner.String, Value: "hello\nworld", Pos: Position{Line: 1, Column: 1}})
	lexer = LexString(`'\U00008a9e'`)
	assert.Equal(t, lexer.Next(), Token{Type: scanner.Char, Value: "\U00008a9e", Pos: Position{Line: 1, Column: 1}})
}

func TestLexBacktickString(t *testing.T) {
	lexer := LexString("`hello\\nworld`")
	assert.Equal(t, lexer.Next(), Token{Type: scanner.String, Value: "hello\\nworld", Pos: Position{Line: 1, Column: 1}})
}

func BenchmarkTextScannerLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lex := TextScannerLexer.Lex(strings.NewReader("hello world 123 hello world 123"))
		ConsumeAll(lex)
	}
}
