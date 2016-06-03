# pg_query_go [![Build Status](https://travis-ci.org/lfittl/pg_query_go.svg)](https://travis-ci.org/lfittl/pg_query_go) [![GoDoc](https://godoc.org/github.com/lfittl/pg_query_go?status.svg)](https://godoc.org/github.com/lfittl/pg_query_go)

Go version of https://github.com/lfittl/pg_query

This Go library and its cgo extension use the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree.

Note that the original Ruby version of this library is much more feature complete.

You can find further background to why a query's parse tree is useful here: https://pganalyze.com/blog/parse-postgresql-queries-in-ruby.html


## Installation

```
go get github.com/lfittl/pg_query_go
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
  "github.com/lfittl/pg_query_go"
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
$ go run main.go
[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": {"name": null, "indirection": null, "val": {"A_CONST": {"val": 1, "location": 7}}, "location": 7}}], "fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]
```

### Parsing a query into Go structs

When working with the query information inside Go its recommended you use the `Parse()` method which returns Go structs:

```go
package main

import (
  "fmt"
  "reflect"
  "github.com/lfittl/pg_query_go"
  nodes "github.com/lfittl/pg_query_go/nodes"
)

func main() {
  tree, err := pg_query.Parse("SELECT 1")
  if err != nil {
    panic(err);
  }

  fmt.Printf("%s\n", reflect.DeepEqual(tree, pg_query.ParsetreeList{
    Statements: []nodes.Node{
      nodes.SelectStmt{
        TargetList: []nodes.Node{
          nodes.ResTarget{
            Val: nodes.A_Const{
              Type: "integer",
              Val: nodes.Value{
                Type: nodes.T_Integer,
                Ival: 1,
              },
              Location: 7,
            },
            Location: 7,
          },
        },
      },
    },
  }));
}
```

You can find all the node struct types in the `nodes/` directory.

## Benchmarks

As it stands, parsing has considerable overhead for complex queries, due to the use of JSON to pass structs across the C <=> Go barrier.

```
BenchmarkParseSelect1-4              	   30000	     43407 ns/op	   14608 B/op	     226 allocs/op
BenchmarkParseSelect2-4              	   10000	    177160 ns/op	   49105 B/op	     742 allocs/op
BenchmarkParseCreateTable-4          	    2000	    573889 ns/op	  149827 B/op	    2123 allocs/op
BenchmarkParseSelect1Parallel-4      	  100000	     16816 ns/op	   14608 B/op	     226 allocs/op
BenchmarkParseSelect2Parallel-4      	   20000	     76318 ns/op	   49105 B/op	     742 allocs/op
BenchmarkParseCreateTableParallel-4  	   10000	    198530 ns/op	  149828 B/op	    2123 allocs/op
```

A good portion of this is due to JSON parsing inside Go so we can work with Go structs - just the raw parser is 10x faster:

```
BenchmarkRawParseSelect1-4           	  500000	      3727 ns/op	     192 B/op	       2 allocs/op
BenchmarkRawParseSelect2-4           	  200000	      9947 ns/op	     672 B/op	       2 allocs/op
BenchmarkRawParseCreateTable-4       	   50000	     28536 ns/op	    2080 B/op	       2 allocs/op
```

Similarly, for query fingerprinting, you might want to use `pg_query.FastFingerprint` to let the C extension handle it:

```
BenchmarkFingerprintSelect1-4        	   30000	     51022 ns/op	   15180 B/op	     244 allocs/op
BenchmarkFingerprintSelect2-4        	   10000	    211596 ns/op	   51919 B/op	     826 allocs/op
BenchmarkFingerprintCreateTable-4    	    2000	    643965 ns/op	  161483 B/op	    2358 allocs/op
BenchmarkFastFingerprintSelect1-4    	  300000	      4442 ns/op	      80 B/op	       2 allocs/op
BenchmarkFastFingerprintSelect2-4    	  200000	      8328 ns/op	      80 B/op	       2 allocs/op
BenchmarkFastFingerprintCreateTable-4	  100000	     22475 ns/op	      80 B/op	       2 allocs/op
```

Normalization is already handled in the C extension, doesn't depend on JSON parsing at all, and is fast:

```
BenchmarkNormalizeSelect1-4          	 1000000	      2121 ns/op	      24 B/op	       2 allocs/op
BenchmarkNormalizeSelect2-4          	  300000	      4219 ns/op	      64 B/op	       2 allocs/op
BenchmarkNormalizeCreateTable-4      	  200000	      7446 ns/op	     144 B/op	       2 allocs/op
```

See `benchmark_test.go` for the queries.

Benchmark numbers from running on a 3.2 GHz Intel Core i5 CPU, OSX 10.11.


## Authors

- [Lukas Fittl](mailto:lukas@fittl.com)


## License

Copyright (c) 2015, Lukas Fittl <lukas@fittl.com><br>
pg_query_go is licensed under the 3-clause BSD license, see LICENSE file for details.

This project includes code derived from the [PostgreSQL project](http://www.postgresql.org/),
see LICENSE.POSTGRESQL for details.
