//go:build required
// +build required

// package gokeep prevents go tooling from stripping the c dependencies.
package gokeep

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/port/win32/arpa"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/port/win32/netinet"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/port/win32/sys"
)
