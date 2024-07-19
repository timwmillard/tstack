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
	"testing"

	"github.com/matryer/is"
)

func Test_split(t *testing.T) {
	is := is.New(t)
	for _, tt := range []struct {
		sql      string
		expected []string
	}{
		{
			sql:      ``,
			expected: []string{``},
		},
		{
			sql:      `select 42`,
			expected: []string{`select 42`},
		},
		{
			sql:      `select $1`,
			expected: []string{`select $1`},
		},
		{
			sql:      `select 42; select 7;`,
			expected: []string{`select 42;`, `select 7;`},
		},
		{
			sql:      `select 42; select 7`,
			expected: []string{`select 42;`, `select 7`},
		},
		{
			sql:      `select 42, ';' ";"; select 7;`,
			expected: []string{`select 42, ';' ";";`, `select 7;`},
		},
		{
			sql: `select * -- foo ; bar
from users -- ; single line comments
where id = $1;
select 1;
`,
			expected: []string{`select * -- foo ; bar
from users -- ; single line comments
where id = $1;`,
				`select 1;`},
		},
		{
			sql: `select * /* ; multi line ;
;
*/
/* /* ; with nesting ; */ */
from users
where id = $1;
select 1;
`,
			expected: []string{`select * /* ; multi line ;
;
*/
/* /* ; with nesting ; */ */
from users
where id = $1;`,
				`select 1;`},
		},
	} {
		actual := split(tt.sql)
		is.Equal(actual, tt.expected)
	}
}
