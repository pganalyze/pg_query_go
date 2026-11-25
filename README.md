# pg_query_go [![GoDoc](https://godoc.org/github.com/pganalyze/pg_query_go/v18?status.svg)](https://godoc.org/github.com/pganalyze/pg_query_go/v18)

Go version of https://github.com/pganalyze/pg_query

This Go library and its cgo extension use the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree.

You can find further background to why a query's parse tree is useful here: https://pganalyze.com/blog/parse-postgresql-queries-in-ruby.html


## Installation

```
go get github.com/pganalyze/pg_query_go/v18@latest
```

Due to compiling parts of PostgreSQL, the first time you build against this library it will take a bit longer.

Expect up to 3 minutes. You can use `go build -x` to see the progress.

## Usage

### Parsing a query into JSON

Put the following in a new Go package, after having installed pg_query as above:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v18"
)

func main() {
	tree, err := pg_query.ParseToJSON("SELECT 1", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", tree)
}
```

Running will output the query's parse tree as JSON:

```json
{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"ival":{"ival":1},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}
```

### Parsing a query into Go structs

When working with the query information inside Go its recommended you use the `Parse()` method which returns Go structs:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v18"
)

func main() {
	result, err := pg_query.Parse("SELECT 42", 0)
	if err != nil {
		panic(err)
	}

	// This will output "42"
	fmt.Printf("%d\n", result.Stmts[0].Stmt.GetSelectStmt().GetTargetList()[0].GetResTarget().GetVal().GetAConst().GetIval().Ival)
}
```

You can find all the node types in the `pg_query.pb.go` Protobuf definition.

### Parser options

`Parse`, `ParseToJSON`, `Fingerprint`, `FingerprintToUInt64` and `Summary` accept parser options
that control how the input is parsed. Pass `0` (`pg_query.ParseDefault`) for the standard
top-level SQL grammar, as shown in the examples above.

The options consist of a parse mode (which grammar to use), optionally combined with flags
that mirror Postgres settings:

```go
// Parse a type name instead of a full SQL statement
result, err := pg_query.Parse("character varying(32)", pg_query.ParseTypeName)

// Parse like a server that has standard_conforming_strings = off
result, err = pg_query.Parse(`SELECT 'a\'b'`, pg_query.DisableStandardConformingStrings)

// Mode and flags can be combined
result, err = pg_query.Parse("integer", pg_query.ParseTypeName|pg_query.DisableEscapeStringWarning)
```

| `ParserOptions`                    | effect                                                                    |
| ---------------------------------- | ------------------------------------------------------------------------- |
| `ParseDefault` (0)                 | parse using the standard top-level SQL grammar                            |
| `ParseTypeName`                    | parse a type name (e.g. `integer`, `character varying(32)`)               |
| `ParsePlPgSqlExpr`                 | parse a PL/pgSQL expression                                               |
| `ParsePlPgSqlAssign1`              | parse a PL/pgSQL assignment (target 1)                                    |
| `ParsePlPgSqlAssign2`              | parse a PL/pgSQL assignment (target 2)                                    |
| `ParsePlPgSqlAssign3`              | parse a PL/pgSQL assignment (target 3)                                    |
| `DisableBackslashQuote`            | `backslash_quote = off` (default is `safe_encoding`, which is effectively on) |
| `DisableStandardConformingStrings` | `standard_conforming_strings = off` (default is on)                       |
| `DisableEscapeStringWarning`       | `escape_string_warning = off` (default is on)                             |

### Deparsing a parse tree back into a SQL statement

In order to go back from a parse tree to a SQL statement, you can use the deparsing functionality:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v18"
)

func main() {
	result, err := pg_query.Parse("SELECT 42", 0)
	if err != nil {
		panic(err)
	}

	result.Stmts[0].Stmt.GetSelectStmt().GetTargetList()[0].GetResTarget().Val = pg_query.MakeAConstStrNode("Hello World", -1)

	stmt, err := pg_query.Deparse(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", stmt)
}
```

This will output the following:

```
SELECT 'Hello World'
```

Note that it is currently not recommended to pass unsanitized input to the deparser, as it may lead to crashes.

### Fingerprinting a query

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v18"
)

func main() {
	fingerprint, err := pg_query.Fingerprint("SELECT * FROM contacts.person WHERE id IN (1, 2, 3, 4);", 0, 0)
	if err != nil {
		panic(err)
	}

	// This will output "5735f5c64dd9f68e"
	fmt.Printf("%s\n", fingerprint)
}
```

Fingerprinting allows you to identify similar queries that are different only
because of the specific object being queried (e.g. different object ids in the
`WHERE` clause) or because of formatting. It is intended to be a superset of
the Postgres `queryid` mechanism: constant values are ignored, so
`... WHERE id = 123` and `... WHERE id = 456` get the same fingerprint.

`Fingerprint` returns the fingerprint as a hex string, use `FingerprintToUInt64`
to get it as a `uint64` instead. Both take the parser options as their second
argument (see above), and the fingerprint options as their third argument.
Pass `0` for either when no specific options are needed.

Postgres 18 significantly changed how schema names and table aliases in `FROM`
clauses are handled for `queryid`, and pg_query follows that by default:

- **Postgres 17 and earlier**: the schema name was significant, and table
  aliases were always ignored.
- **Postgres 18 (default)**: schema names are ignored, and a table alias (if
  present) is used instead of the relation name.

This means that fingerprints differ from those calculated by pg_query_go v6
and earlier, unless you pass `pg_query.FingerprintRangeVarPG17Compat`:

```go
// Postgres 18 behaviour (default): schema is ignored and the alias replaces the relation name
a, _ := pg_query.Fingerprint("SELECT * FROM public.users u", 0, 0)
b, _ := pg_query.Fingerprint("SELECT * FROM myschema.users u", 0, 0)
// a == b ("6640d8be64880eed")

// Postgres 17 compatible fingerprints take the schema into account
a, _ = pg_query.Fingerprint("SELECT * FROM public.users u", 0, pg_query.FingerprintRangeVarPG17Compat)
b, _ = pg_query.Fingerprint("SELECT * FROM myschema.users u", 0, pg_query.FingerprintRangeVarPG17Compat)
// a != b ("d3198777453aef12" vs "91e880cf13f4f219")
```

Available fingerprint options (these can be combined with a bitwise OR):

| `FingerprintOptions`               | effect                                                                    |
| ---------------------------------- | ------------------------------------------------------------------------- |
| `FingerprintDefault` (0)           | Postgres 18 query ID behaviour                                            |
| `FingerprintRangeVarIgnoreAliases` | relation names are always fingerprinted; aliases are ignored              |
| `FingerprintRangeVarIncludeSchema` | schema names are also fingerprinted in SELECT/DML statements              |
| `FingerprintRangeVarPG17Compat`    | combination of the two above, matching Postgres 17 (and pg_query_go v6) fingerprints |
| `FingerprintFullRelname`           | fingerprint the full relation name instead of ignoring 2+ consecutive digits (which by default groups date/number-suffixed partitions together) |

### Parsing a PL/pgSQL function into JSON (Experimental)

Put the following in a new Go package, after having installed pg_query as above:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v18"
)

func main() {
	tree, err := pg_query.ParsePlPgSqlToJSON(
		`CREATE OR REPLACE FUNCTION cs_fmt_browser_version(v_name varchar, v_version varchar)
  			RETURNS varchar AS $$
  			BEGIN
  			    IF v_version IS NULL THEN
  			        RETURN v_name;
  			    END IF;
  			    RETURN v_name || '/' || v_version;
  			END;
  			$$ LANGUAGE plpgsql;`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", tree)
}
```

Running will output the functions's parse tree as JSON:

```json
$ go run main.go
[
{"PLpgSQL_function":{"datums":[{"PLpgSQL_var":{"refname":"v_name","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}},{"PLpgSQL_var":{"refname":"v_version","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}},{"PLpgSQL_var":{"refname":"found","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}}],"action":{"PLpgSQL_stmt_block":{"lineno":2,"body":[{"PLpgSQL_stmt_if":{"lineno":3,"cond":{"PLpgSQL_expr":{"query":"v_version IS NULL"}},"then_body":[{"PLpgSQL_stmt_return":{"lineno":4,"expr":{"PLpgSQL_expr":{"query":"v_name"}}}}]}},{"PLpgSQL_stmt_return":{"lineno":6,"expr":{"PLpgSQL_expr":{"query":"v_name || '/' || v_version"}}}}]}}}}
]
```

## Benchmarks

```
$ make benchmark
go build -a
go test -test.bench=. -test.run=XXX -test.benchtime 10s -test.benchmem -test.cpu=4
goos: darwin
goarch: arm64
pkg: github.com/pganalyze/pg_query_go/v18
BenchmarkParseSelect1-4                          2874156              4186 ns/op            1040 B/op         18 allocs/op
BenchmarkParseSelect2-4                           824781             14572 ns/op            2832 B/op         57 allocs/op
BenchmarkParseCreateTable-4                       351037             34591 ns/op            8480 B/op        149 allocs/op
BenchmarkParseSelect1Parallel-4                  9027080              1320 ns/op            1040 B/op         18 allocs/op
BenchmarkParseSelect2Parallel-4                  2745390              4369 ns/op            2832 B/op         57 allocs/op
BenchmarkParseCreateTableParallel-4              1000000             10487 ns/op            8480 B/op        149 allocs/op
BenchmarkRawParseSelect1-4                       3778771              3183 ns/op             128 B/op          3 allocs/op
BenchmarkRawParseSelect2-4                       1000000             10985 ns/op             288 B/op          3 allocs/op
BenchmarkRawParseCreateTable-4                    460714             26397 ns/op            1056 B/op          3 allocs/op
BenchmarkRawParseSelect1Parallel-4              13338790               902.7 ns/op           128 B/op          3 allocs/op
BenchmarkRawParseSelect2Parallel-4               4060762              2956 ns/op             288 B/op          3 allocs/op
BenchmarkRawParseCreateTableParallel-4           1709883              7001 ns/op            1056 B/op          3 allocs/op
BenchmarkFingerprintSelect1-4                    6394882              1875 ns/op              48 B/op          2 allocs/op
BenchmarkFingerprintSelect2-4                    2865390              4174 ns/op              48 B/op          2 allocs/op
BenchmarkFingerprintCreateTable-4                1688920              7143 ns/op              48 B/op          2 allocs/op
BenchmarkNormalizeSelect1-4                     10604962              1133 ns/op              32 B/op          2 allocs/op
BenchmarkNormalizeSelect2-4                      6226136              1938 ns/op              64 B/op          2 allocs/op
BenchmarkNormalizeCreateTable-4                  4542387              2635 ns/op             144 B/op          2 allocs/op
PASS
ok      github.com/pganalyze/pg_query_go/v18     258.376s

```

Note that allocation counts exclude the cgo portion, so they are higher than shown here.

See `benchmark_test.go` for details on the benchmarks.


## Authors

- [Lukas Fittl](mailto:lukas@fittl.com)


## License

Copyright (c) 2015, Lukas Fittl <lukas@fittl.com><br>
Copyright (c) 2016-2025, Duboce Labs, Inc. (pganalyze) <team@pganalyze.com>
pg_query_go is licensed under the 3-clause BSD license, see LICENSE file for details.

This project includes code derived from the [PostgreSQL project](http://www.postgresql.org/),
see LICENSE.POSTGRESQL for details.
