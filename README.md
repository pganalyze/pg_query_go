# pg_query_go

This is a fork of the original [pg_query_go](https://github.com/pganalyze/pg_query_go) library.

The fork widens access to the deparser to allow deparsing of certain other node types, rather than just full SQL statements.

It is used as part of the `sql2pgroll` package in `pgroll`.

The intention is to get these changes merged upstream so that this fork can be removed.
