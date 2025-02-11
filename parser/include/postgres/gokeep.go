//go:build required
// +build required

// package gokeep prevents go tooling from stripping the c dependencies.
package gokeep

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/access"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/archive"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/catalog"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/commands"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/common"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/datatype"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/executor"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/foreign"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/jit"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/lib"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/libpq"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/mb"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/nodes"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/optimizer"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/parser"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/partitioning"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/port"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/portability"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/postmaster"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/regex"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/replication"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/rewrite"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/storage"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/tcop"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/tsearch"
	_ "github.com/pganalyze/pg_query_go/v6/parser/include/postgres/utils"
)
