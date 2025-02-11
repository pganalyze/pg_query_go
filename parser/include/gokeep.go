//go:build required
// +build required

// package gokeep prevents go tooling from stripping the c dependencies.
package gokeep

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/protobuf"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/protobuf-c"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/xxhash"
)
