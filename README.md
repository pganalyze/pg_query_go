# pg_query_go [![GoDoc](https://godoc.org/github.com/pganalyze/pg_query_go?status.svg)](https://godoc.org/github.com/pganalyze/pg_query_go)

Go version of https://github.com/pganalyze/pg_query

This Go library and its cgo extension use the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree.

You can find further background to why a query's parse tree is useful here: https://pganalyze.com/blog/parse-postgresql-queries-in-ruby.html


## Installation

```
go get github.com/pganalyze/pg_query_go/v2
```

Due to compiling parts of PostgreSQL, the first time you build against this library it will take a bit longer.

Expect up to 3 minutes. You can use `go build -x` to see the progress.

## Usage with Go modules

When integrating this library using Go modules, and using a vendor/ directory,
you will need to explicitly copy over some of the C build files, since Go does
not copy files in subfolders without .go files whilst vendoring.

The best way to do so is to use [modvendor](https://github.com/goware/modvendor),
and vendor your modules like this:

```
go mod vendor
go get -u github.com/goware/modvendor
modvendor -copy="**/*.c **/*.h **/*.proto" -v
```

## Usage

### Parsing a query into JSON

Put the following in a new Go package, after having installed pg_query as above:

```go
package main

import (
  "fmt"
  "github.com/pganalyze/pg_query_go"
)

func main() {
  tree, err := pg_query.ParseToJSON("SELECT 1")
  if err != nil {
    panic(err);
  }
  fmt.Printf("%s\n", tree)
}
```

Running will output the query's parse tree as JSON:

```json
{"version":130002,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"val":{"Integer":{"ival":1}},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}
```

### Parsing a query into Go structs

When working with the query information inside Go its recommended you use the `Parse()` method which returns Go structs:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v2"
)

func main() {
	result, err := pg_query.Parse("SELECT 42")
	if err != nil {
		panic(err)
	}

	// This will output "42"
	fmt.Printf("%d\n", result.Stmts[0].Stmt.GetSelectStmt().GetTargetList()[0].GetResTarget().GetVal().GetAConst().GetVal().GetInteger().Ival)
}
```

You can find all the node types in the `pg_query.pb.go` Protobuf definition.

### Deparsing a parse tree back into a SQL statement (Experimental)

In order to go back from a parse tree to a SQL statement, you can use the deparsing functionality:

```go
package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v2"
)

func main() {
	result, err := pg_query.Parse("SELECT 42")
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

### Parsing a PL/pgSQL function into JSON (Experimental)

Put the following in a new Go package, after having installed pg_query as above:

```go
package main

import (
  "fmt"
  "github.com/pganalyze/pg_query_go/v2"
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
    panic(err);
  }
  fmt.Printf("%s\n", tree)
}
```

Running will output the functions's parse tree as JSON:

```json
$ go run main.go
[
{"PLpgSQL_function":{"datums":[{"PLpgSQL_var":{"refname":"found","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}}],"action":{"PLpgSQL_stmt_block":{"lineno":2,"body":[{"PLpgSQL_stmt_if":{"lineno":3,"cond":{"PLpgSQL_expr":{"query":"SELECT v_version IS NULL"}},"then_body":[{"PLpgSQL_stmt_return":{"lineno":4,"expr":{"PLpgSQL_expr":{"query":"SELECT v_name"}}}}]}},{"PLpgSQL_stmt_return":{"lineno":6,"expr":{"PLpgSQL_expr":{"query":"SELECT v_name || '/' || v_version"}}}}]}}}}
]
```

## Benchmarks

```
BenchmarkParseSelect1-4                  	 1542726	      7757 ns/op	    1168 B/op	      21 allocs/op
BenchmarkParseSelect2-4                  	  496621	     24027 ns/op	    3184 B/op	      63 allocs/op
BenchmarkParseCreateTable-4              	  231754	     51624 ns/op	    8832 B/op	     157 allocs/op
BenchmarkParseSelect1Parallel-4          	 5451890	      2213 ns/op	    1168 B/op	      21 allocs/op
BenchmarkParseSelect2Parallel-4          	 1711480	      7067 ns/op	    3184 B/op	      63 allocs/op
BenchmarkParseCreateTableParallel-4      	  759412	     16157 ns/op	    8832 B/op	     157 allocs/op
BenchmarkRawParseSelect1-4               	 2311986	      5158 ns/op	     192 B/op	       5 allocs/op
BenchmarkRawParseSelect2-4               	  721333	     16517 ns/op	     384 B/op	       5 allocs/op
BenchmarkRawParseCreateTable-4           	  328119	     35675 ns/op	    1248 B/op	       5 allocs/op
BenchmarkRawParseSelect1Parallel-4       	 8378274	      1412 ns/op	     192 B/op	       5 allocs/op
BenchmarkRawParseSelect2Parallel-4       	 2650692	      4553 ns/op	     384 B/op	       5 allocs/op
BenchmarkRawParseCreateTableParallel-4   	 1000000	     10335 ns/op	    1248 B/op	       5 allocs/op
BenchmarkFingerprintSelect1-4           	 5012028	      2388 ns/op	     112 B/op	       4 allocs/op
BenchmarkFingerprintSelect2-4           	 2391704	      5026 ns/op	     112 B/op	       4 allocs/op
BenchmarkFingerprintCreateTable-4       	 1304545	      9601 ns/op	     112 B/op	       4 allocs/op
BenchmarkNormalizeSelect1-4              	 8767273	      1360 ns/op	      72 B/op	       4 allocs/op
BenchmarkNormalizeSelect2-4              	 4465364	      2756 ns/op	     104 B/op	       4 allocs/op
BenchmarkNormalizeCreateTable-4          	 2738284	      4345 ns/op	     184 B/op	       4 allocs/op
```

Note that allocation counts exclude the cgo portion, so they are higher than shown here.

See `benchmark_test.go` for details on the benchmarks.


## Authors

- [Lukas Fittl](mailto:lukas@fittl.com)


## License

Copyright (c) 2015, Lukas Fittl <lukas@fittl.com><br>
pg_query_go is licensed under the 3-clause BSD license, see LICENSE file for details.

This project includes code derived from the [PostgreSQL project](http://www.postgresql.org/),
see LICENSE.POSTGRESQL for details.
