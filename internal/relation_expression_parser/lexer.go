package relation_expression_parser

import (
	"unicode"
	"unicode/utf8"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

const eof rune = -1

type token struct {
	Type     tokenType
	Lexeme   string
	StartPos int
	EndPos   int
}

type predicate func(rune) bool

func isIdentifierRune(elem rune) bool {
	// TODO read this https://www.unicode.org/reports/tr31/#Introduction
	return unicode.IsLetter(elem) || unicode.IsDigit(elem) || elem == '_'
}

type tokenType int

const (
	tokenEof tokenType = iota
	tokenIdentifier
	tokenUnion
	tokenDifference
	tokenIntersection
	tokenGroupBegin
	tokenGroupEnd
	tokenArrow
	tokenThis
	tokenError
)

type Lexer struct {
	input     string
	startPos  int
	currPos   int
	lastWidth int
	q         chan token
}

func NewLexer(input string) (Lexer, <-chan token) {
	q := make(chan token, 10)
	l := Lexer{
		input,
		0,
		0,
		0,
		q,
	}
	return l, q
}

// return next rune as a string
func (l *Lexer) peek() (rune, int) {
	if l.currPos == len(l.input) {
		return eof, 0
	}

	return utf8.DecodeRuneInString(l.input[l.currPos:])
}

// return one item in the lexer
// should prob panic if it steps past startPos
func (l *Lexer) backtrack() {
	l.currPos -= l.lastWidth
}

// proceeds lexer to the next rune in the input string
func (l *Lexer) step() rune {
	char, width := l.peek()
	l.currPos += width
	l.lastWidth = width
	return char
}

func (l *Lexer) reset() {
	l.currPos = l.startPos
}

func (l *Lexer) ignore() {
	l.startPos = l.currPos
}

// step one position in parser if predicate is true
func (l *Lexer) stepIf(f predicate) bool {
	r, _ := l.peek()
	if f(r) {
		l.step()
		return true
	}
	return false
}

// step one position in parser if predicate is true
func (l *Lexer) stepIfRune(r rune) bool {
	nextRune, _ := l.peek()
	if nextRune == r {
		l.step()
		return true
	}
	return false
}

// move lexer while predicate is true
func (l *Lexer) stepWhile(f predicate) int {
	i := 0
	for ; l.stepIf(f); i++ {
	}
	return i
}

// attempts to consume string from input buffer
// if it finds string, steps lexer and returns true
// otherwise, keeps original window and return false
func (l *Lexer) consumeString(str string) bool {
	oldStart := l.currPos
	ok := true
	for _, r := range str {
		ok = l.stepIfRune(r)
		if !ok {
			break
		}
	}
	// thsi could fail due to an eof

	if !ok {
		l.currPos = oldStart
	}
	return ok
}

// move lexer's starPos to the first non-whitespace character
func (l *Lexer) skipSpaces() {
	l.stepWhile(unicode.IsSpace)
	l.ignore()
}

func (l *Lexer) lexTerminal(typ tokenType, str string) bool {
	ok := l.consumeString(str)
	if ok {
		l.emit(typ)
	}

	return ok
}

func (l *Lexer) lexEOF() bool {
	char, _ := l.peek()
	return char == eof
}

func (l *Lexer) lexIdentifier() bool {
	count := l.stepWhile(isIdentifierRune)
	ok := count > 0
	if ok {
		l.emit(tokenIdentifier)
	}

	return ok
}

func (l *Lexer) scan() {
	// consume as many operators, expressions or parenthesis
	// as it can
	for {
		if l.lexTerminal(tokenGroupBegin, domain.GroupBeginLexeme) {
		} else if l.lexTerminal(tokenGroupEnd, domain.GroupEndLexeme) {
		} else if l.lexTerminal(tokenArrow, domain.ArrowLexeme) { // arrow has higher precedence than difference
		} else if l.lexTerminal(tokenUnion, domain.UnionLexeme) {
		} else if l.lexTerminal(tokenDifference, domain.DifferenceLexeme) {
		} else if l.lexTerminal(tokenIntersection, domain.IntersectionLexeme) {
		} else if l.lexTerminal(tokenThis, domain.ThisLexeme) {
		} else if l.lexIdentifier() {
		} else if char, _ := l.peek(); unicode.IsSpace(char) {
			l.skipSpaces()
		} else if l.lexEOF() {
			l.emit(tokenEof)
			break
		} else {
			l.emitError("unknown token")
			break
		}
	}
}

func (l *Lexer) emit(tt tokenType) {
	lexeme := l.input[l.startPos:l.currPos]
	t := token{
		Type:     tt,
		Lexeme:   lexeme,
		StartPos: l.startPos,
		EndPos:   l.currPos,
	}
	l.q <- t
	l.ignore()
}

func (l *Lexer) emitError(msg string) {
	msg = msg + ": context: " + l.input[l.startPos:]
	t := token{
		Type:     tokenError,
		Lexeme:   msg,
		StartPos: l.startPos,
		EndPos:   l.currPos,
	}
	l.q <- t

	l.ignore()
}

func (l *Lexer) Lex() {
	l.scan()
	close(l.q)
}
