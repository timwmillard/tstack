# sqltest

Snapshot database output and assert database tables.


## Usage

See example below.

Use the `-update` flag when running `go test` to update the result snapshots.

## Example

`users_test.go`

```go
func TestCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestCreateUser")
	}
	t.Parallel()

	ctx := context.Background()
	db := createDatabase(ctx, t)

    user := User{
        FirstName: "Tim",
        LastName: "Millard",
    }

	_, err := CreateUser(ctx, db, user)
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	sqltest := sqltest.New(t, db)
	sqltest.AssertQuery(ctx, "testdata/createuser.sql", user.ID.String())
}
```

`testdata/createuser.sql`

```sql
select first_name, last_name
from users;
```

`testdata/createuser.results`

```
-- sqltest
select first_name, last_name
from users;
+------------+-----------+
| first_name | last_name |
+------------+-----------+
| Tim        | Millard   |
+------------+-----------+
```
