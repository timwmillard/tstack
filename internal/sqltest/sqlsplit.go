package sqltest

// Copyright (c) 2011-2020 Jack Christensen
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

import (
	"strings"
	"unicode/utf8"
)

// split splits sql into into a slice of strings each containing one SQL statement.
func split(sql string) []string {
	l := &sqlLexer{
		src:     sql,
		stateFn: rawState,
	}

	for l.stateFn != nil {
		l.stateFn = l.stateFn(l)
	}

	if len(l.statements) == 0 {
		l.statements = []string{sql}
	}

	return l.statements
}

type sqlLexer struct {
	src     string
	start   int
	pos     int
	nested  int // multiline comment nesting level.
	stateFn stateFn

	statements []string
}

func (l *sqlLexer) addStatement(s string) {
	s = strings.TrimSpace(s)
	if len(s) > 0 {
		l.statements = append(l.statements, s)
	}
}

type stateFn func(*sqlLexer) stateFn

func rawState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case 'e', 'E':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune == '\'' {
				l.pos += width
				return escapeStringState
			}
		case '\'':
			return singleQuoteState
		case '"':
			return doubleQuoteState
		case ';':
			l.addStatement(l.src[l.start:l.pos])
			l.start = l.pos
			return rawState
		case '-':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune == '-' {
				l.pos += width
				return oneLineCommentState
			}
		case '/':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune == '*' {
				l.pos += width
				return multilineCommentState
			}
		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}

func singleQuoteState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case '\'':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune != '\'' {
				return rawState
			}
			l.pos += width
		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}

func doubleQuoteState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case '"':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune != '"' {
				return rawState
			}
			l.pos += width
		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}

func escapeStringState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case '\\':
			_, width = utf8.DecodeRuneInString(l.src[l.pos:])
			l.pos += width
		case '\'':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune != '\'' {
				return rawState
			}
			l.pos += width
		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}

func oneLineCommentState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case '\\':
			_, width = utf8.DecodeRuneInString(l.src[l.pos:])
			l.pos += width
		case '\n', '\r':
			return rawState
		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}

func multilineCommentState(l *sqlLexer) stateFn {
	for {
		r, width := utf8.DecodeRuneInString(l.src[l.pos:])
		l.pos += width

		switch r {
		case '/':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune == '*' {
				l.pos += width
				l.nested++
			}
		case '*':
			nextRune, width := utf8.DecodeRuneInString(l.src[l.pos:])
			if nextRune != '/' {
				continue
			}

			l.pos += width
			if l.nested == 0 {
				return rawState
			}
			l.nested--

		case utf8.RuneError:
			if l.pos-l.start > 0 {
				l.addStatement(l.src[l.start:l.pos])
				l.start = l.pos
			}
			return nil
		}
	}
}
