# pg_query.go [![Build Status](https://travis-ci.org/lfittl/pg_query.go.svg)](https://travis-ci.org/lfittl/pg_query.go)

Experimental Go version of https://github.com/lfittl/pg_query

This Go library and its cgo extension use the actual PostgreSQL server source to parse SQL queries and return the internal PostgreSQL parse tree.

Note that the original Ruby version of this library is much more feature complete.

You can find further background to why a query's parse tree is useful here: https://pganalyze.com/blog/parse-postgresql-queries-in-ruby.html


## Installation

```
cd $GOPATH/src
git clone git://github.com/lfittl/pg_query.go.git github.com/lfittl/pg_query.go
cd github.com/lfittl/pg_query.go
make
go build
```

Due to compiling parts of PostgreSQL, running `make` will take a while. Expect up to 5 minutes.


## Usage

### Parsing a query

Put the following in a new Go package, after having installed pg_query as above:

```go
package main

import (
  "fmt"
  "github.com/lfittl/pg_query.go"
)

func main() {
  tree := pg_query.Parse("SELECT 1")
  fmt.Printf("%s\n", tree)
}
```

Running will output the query's parse tree:

```json
$ go run main.go
[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": {"name": null, "indirection": null, "val": {"A_CONST": {"val": 1, "location": 7}}, "location": 7}}], "fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]
```


## Authors

- [Lukas Fittl](mailto:lukas@fittl.com)


## License

Copyright (c) 2015, Lukas Fittl <lukas@fittl.com><br>
pg_query.go is licensed under the 3-clause BSD license, see LICENSE file for details.

Query normalization code:<br>
Copyright (c) 2008-2015, PostgreSQL Global Development Group
