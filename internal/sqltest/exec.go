package sqltest

import (
	"context"
	"os"
)

func (test Tester) ExecQuery(ctx context.Context, sqlFilename string, args ...any) {
	test.t.Helper()

	query, err := os.ReadFile(sqlFilename)
	if err != nil {
		test.t.Fatalf("Unable to read file %s: %v", sqlFilename, err)
	}

	queries := split(string(query))
	for _, q := range queries {
		rows, err := test.db.Query(ctx, q, args...)
		if err != nil {
			test.t.Fatalf("Unable to exec query: %v", err)
		}
		defer rows.Close()
	}
}
