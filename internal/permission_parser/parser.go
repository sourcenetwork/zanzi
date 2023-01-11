package permission_parser

import (
	"fmt"
)

// parser function type
// receives slice of tokens, return some type, error and another token slice
type parserFn[T any] func([]token) (T, error, []token)

// unit parser
// consumes any item from token slice, fails if slice is empty
func unit(tokens []token) (token, error, []token) {
	if len(tokens) == 0 {
		return token{}, fmt.Errorf("failed parsing: no tokens left"), tokens
	}

	return tokens[0], nil, tokens[1:]
}

func buildErr(tokens []token, msg string, args ...any) error {
	msg = fmt.Sprintf(msg, args...)
	return fmt.Errorf("error parsing: %v: have: %v", msg, tokens)
}

// combinators

// many generates a new parser that executes as many times as possible
// many never fails, even if it does not produce anythign
func many[T any](parser parserFn[T]) parserFn[[]T] {
	return func(tokens []token) ([]T, error, []token) {
		if len(tokens) == 0 {
			return nil, nil, tokens
		}

		var tail []token = tokens
		var ts []T
		for {
			t, err, newTail := parser(tail)
			if err != nil {
				return ts, nil, tail
			}
			tail = newTail
			ts = append(ts, t)
		}
	}
}

// first returns a parser that tries parsers until the first sucessful parsing
func first[T any](errMsg string, parsers ...parserFn[T]) parserFn[T] {
	return func(tokens []token) (T, error, []token) {
		var zero T
		if len(tokens) == 0 {
			return zero, buildErr(tokens, errMsg), tokens
		}

		for _, parser := range parsers {
			t, err, sliced := parser(tokens)
			if err == nil {
				return t, nil, sliced
			}
		}

		return zero, buildErr(tokens, errMsg), tokens
	}
}

func parseIf[T any](predicate func(token) bool, errFmt string, args ...any) parserFn[token] {
	return func(tokens []token) (token, error, []token) {
		t, err, sliced := unit(tokens)
		if err != nil {
			return token{}, err, sliced
		}
		if predicate(t) {
			return t, nil, sliced
		}
		return token{}, buildErr(tokens, errFmt, args), tokens
	}
}

// parseIfType return next token if it is of given type
func parseIfType(ttype tokenType) parserFn[token] {
	predicate := func(t token) bool { return t.Type == ttype }
	return parseIf[token](predicate, "want token of type %v", ttype)
}

// parseIfAnyType return next token if it is of given type
func parseIfAnyType(types ...tokenType) parserFn[token] {
	predicate := func(t token) bool {
		for _, wantedType := range types {
			if t.Type == wantedType {
				return true
			}
		}
		return false
	}
	return parseIf[token](predicate, "possible accepted tokens %v", types)
}
