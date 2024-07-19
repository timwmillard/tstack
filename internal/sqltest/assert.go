// Package sqltest allows for test queries to be compared.
//
// Inspired by:
// Go Testing By Example (GopherCon Australia 2023) by Russ Cox
// https://www.youtube.com/watch?v=X4rxi9jStLo

package sqltest

import (
	"bytes"
	"context"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/jackc/pgx/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var update = flag.Bool("update", false, "update sql results files with output")
var preview = flag.Bool("preview", false, "preview sql results to stdout")

type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type Tester struct {
	t  *testing.T
	db Querier
}

func New(t *testing.T, db Querier) Tester {
	return Tester{
		t:  t,
		db: db,
	}
}

func (test Tester) AssertQuery(ctx context.Context, sqlFilename string, args ...string) {
	test.t.Helper()

	name := strings.TrimSuffix(sqlFilename, filepath.Ext(sqlFilename))
	resultsFilename := name + ".results"

	queries, err := os.ReadFile(sqlFilename)
	if err != nil {
		test.t.Fatalf("Unable to read file %s: %v", sqlFilename, err)
	}

	if *preview {
		err = PrintQueries(ctx, os.Stderr, test.db, string(queries), args...)
		if err != nil {
			test.t.Fatalf("Unable to print query: %v", err)
		}
	}

	if *update {
		file, err := os.Create(resultsFilename)
		if err != nil {
			test.t.Fatalf("Unable to open testdata file %s: %v", name, err)
		}
		defer file.Close()
		err = PrintQueries(ctx, file, test.db, string(queries), args...)
		if err != nil {
			test.t.Fatalf("Unable to print query: %v", err)
		}
		test.t.Logf("Updated %s with query output", resultsFilename)

		return
	}

	want, err := os.ReadFile(resultsFilename)
	if err != nil {
		test.t.Fatalf("Unable to open results file [hint: run tests with -update]: %v", err)
	}

	got := &bytes.Buffer{}
	err = PrintQueries(ctx, got, test.db, string(queries), args...)
	if err != nil {
		test.t.Fatalf("unable to print query: %v", err)
	}

	if string(want) != got.String() {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(want), got.String(), false)
		test.t.Errorf("%s mismatch (\x1b[31m-want\x1b[0m \x1b[32m+got\x1b[0m):\n%s",
			resultsFilename,
			dmp.DiffPrettyText(diffs),
		)
	}
}

func PrintQueries(ctx context.Context, wr io.Writer, db Querier, query string, args ...string) error {
	queries := split(query)
	for _, q := range queries {
		err := PrintQuery(ctx, wr, db, q, args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func PrintQuery(ctx context.Context, wr io.Writer, db Querier, query string, args ...string) error {

	var anyArgs []any
	for _, a := range args {
		anyArgs = append(anyArgs, a)
	}

	anyArgs = append([]any{pgx.QueryResultFormats{pgx.TextFormatCode}}, anyArgs...)
	rows, err := db.Query(ctx, query, anyArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return err
	}

	_, err = wr.Write([]byte("-- sqltest\n" + query + "\n"))
	if err != nil {
		return err
	}
	printRows(wr, rows)
	_, err = wr.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}

func printRows(wr io.Writer, rows pgx.Rows) {

	t := table.NewWriter()
	t.SetOutputMirror(wr)
	t.Style().Format.Header = text.FormatDefault

	// Print column headers
	columns := rows.FieldDescriptions()
	header := make(table.Row, 0, len(columns))
	for _, col := range columns {
		header = append(header, col.Name)
	}
	t.AppendHeader(header)

	// Print query rows
	for rows.Next() {
		values := rows.RawValues()

		// Print each row
		row := make(table.Row, 0, len(columns))
		for _, value := range values {
			row = append(row, string(value))
		}
		t.AppendRow(row)
	}
	t.Render()
}
