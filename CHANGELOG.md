# Changelog

## Unreleased

* ...


## 18.1.0     2026-09-30

* Switch to versioning by Postgres major version, matching libpg_query
  - The Go module path is now `github.com/pganalyze/pg_query_go/v18`, update your
    imports accordingly (the previous release was 6.2.5, under `/v6`)
* Upgrade to libpg_query 18.1.0
  - Updates to the Postgres 18 parser (Postgres 18.6)
  - Rework PL/pgSQL parsing to use proper type definitions
    - Instead of hacking the PL/pgSQL parser into accepting our constructs, this
      imitates actual `CREATE FUNCTION` logic and PL/pgSQL compilation for a function
  - Return errors for PL/pgSQL statements without bodies
    - This avoids an assertion failure or crash when `CREATE FUNCTION` or `DO` omits
      its function body
  - Security fix: Heap out-of-bounds write and read in Normalize ([GHSA-6ggm-xmc9-8ffg](https://github.com/pganalyze/libpg_query/security/advisories/GHSA-6ggm-xmc9-8ffg))
    - Same fix as in 6.2.5, see details below
  - Add stack overflow crash protection, error out instead
    - Overly deep queries now return the standard Postgres "stack depth limit exceeded"
      error instead of crashing the process
  - Switch Protobuf implementation from protobuf-c to upb
    - upb is developed as part of the main Protobuf project, and is substantially faster,
      in part due to its built-in arena allocation
    - upb also allows limiting parse depth for complex Protobuf input, avoiding crashes
  - Ignore comments when parsing queries, only treat them as significant for scanning
    - This fixes parse errors when comments are placed between related tokens
      (e.g. `NOT /* comment */ IN`), or between string literals that get concatenated
  - Parser: Avoid quadratic memory use for rules that involve dotted names
  - Fingerprinting:
    - Rework alias/schema name handling to match Postgres 18 query IDs
      - This changes fingerprints compared to prior releases: By default, in SELECT/DML
        statements the alias name replaces the relation name when present, and schema
        names are ignored
      - Use `FingerprintWithOpts` with `FingerprintRangeVarPG17Compat` to get fingerprints
        that match prior releases
    - Ignore `NOTIFY` payloads, similar to channel names
    - Ignore role names (e.g. in `CREATE ROLE`, `DROP ROLE`, `GRANT` and `ALTER ... RENAME`)
    - Include `BEGIN`/`START TRANSACTION` options (e.g. read-only, isolation level)
    - Apply the depth cutoff when recursing into specific node types
      - This changes fingerprints for set operation chains deeper than 100 levels,
        which are now cut off consistently like other deeply nested nodes
  - Deparser:
    - Add strict checking for unexpected pointer values
      - This ensures that a bad input parse tree doesn't cause the deparser to crash, and
        instead returns an error
    - Rework when parentheses are added based on operator precedence
      - This fixes cases where the deparsed SQL changed meaning or was invalid, and
        avoids adding unnecessary parentheses in others
    - Rework comment handling for Postgres 18 multi-statement strings
      - In Postgres 18, comments and whitespace between statements are excluded from the
        statements, the next statement starts at the first non-whitespace (and
        non-comment) character
  - Normalize:
    - Add support for `NOTIFY` statements
    - Avoid undefined behaviour for overly large parameter references
    - Fix handling of `U&` special constants in DefElem nodes
    - Don't swallow stack depth errors and return a partially normalized query
  - Summary:
    - Fix a relation going missing when a CTE shares its name
    - Fix memory leak when the tree walk throws an error
* Add `FingerprintWithOpts` and `FingerprintToUInt64WithOpts` functions
  - These accept a `FingerprintOption` bitmask to control how fingerprints are calculated:
    - `FingerprintRangeVarIgnoreAliases` always fingerprints relation names and ignores aliases
    - `FingerprintRangeVarIncludeSchema` also fingerprints schema names in SELECT/DML statements
    - `FingerprintRangeVarPG17Compat` combines both, matching how Postgres 17 and earlier
      calculate query IDs, and how prior releases calculated fingerprints
    - `FingerprintRelnameFull` fingerprints the full relation name, instead of the default
      behavior of ignoring 2+ consecutive digits (which groups queries on
      date/number-suffixed tables together)
* `MakeNotNullConstraintNode` now sets `IsEnforced` and `InitiallyValid`, matching
  how Postgres 18 represents `NOT NULL` constraints


## 6.2.5     2026-09-30

* Upgrade to libpg_query 17-6.2.5
  - Security fix: Heap out-of-bounds write and read in Normalize ([GHSA-6ggm-xmc9-8ffg](https://github.com/pganalyze/libpg_query/security/advisories/GHSA-6ggm-xmc9-8ffg))
    - When normalizing certain utility statements (e.g. `DO ... LANGUAGE`, statements with
      string options, or `CREATE/ALTER SUBSCRIPTION ... CONNECTION`), Normalize searched
      the query text for the location of string constants, which could yield wrong
      locations for crafted input. This could cause out-of-bounds writes and reads on
      the heap, leaking process memory in the normalized output or crashing the process.
    - Constant locations are now recorded by the parser instead, and the normalizer checks
      at runtime that constant locations never overlap
    - This adds new location fields to the parse tree output (`DefElem.ArgLocation`,
      `NotifyStmt.PayloadLocation`, `CreateSubscriptionStmt.ConninfoLocation` and
      `AlterSubscriptionStmt.ConninfoLocation`). Like other location fields, these are
      ignored for fingerprinting.
    - Applications that normalize untrusted query text should upgrade
    - Reported by Paul Gerste (Cure53)
  - Deparser:
    - Add strict checking for unexpected pointer values
      - This ensures that a bad input parse tree doesn't cause the deparser to crash, and
        instead returns an error
      - Reported by Paul Gerste (Cure53)
    - Preserve parentheses around subscripted array constructors
      - This prevents `(ARRAY[...])[...]` from being deparsed as invalid SQL
    - Fix handling of constraint key named `value` in `ALTER TABLE`
  - Normalize:
    - Add support for `NOTIFY` statements
    - Avoid undefined behaviour for overly large parameter references
    - Fix compiler warning due to mismatch between `int64` and `long`
* Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 [#111](https://github.com/pganalyze/pg_query_go/pull/111)


## 6.2.2     2026-01-27

* Upgrade to libpg_query 17-6.2.2
  - Add IsUtilityStmt function to determine if query text contains utility statements
    - This is a fast check for callers that don't actually need the parse tree itself
  - Add Summary function to get summary information about a statement quickly
    - This allows gathering certain information, for example which tables are referenced in a
      statement, without requiring a Protobuf serialization step in a higher level library
    - Additionally this can also be used to perform "smart truncation" of a query by
      omitting deeply nested information (e.g. a CTE definition, or a target list) whilst
        preserving more essential parts like the FROM clause
  - Normalize:
    - Fix handling of special strings in DefElem
      - This avoids a crash when running the normalize function on certain utility statements
    - Add support for CALL statements
  - Deparser:
    - Introduce pretty printing / formatting
      - Introduces a new optional pretty print mode that emits a human readable
        output. A detailed explanation of the mechanism can be found at the start
        of the deparser file.
    - Rework handling of expressions inside typecasts
      - Prefer (..)::type syntax, unless we are already in a function call.
    - Use lowercase keywords in xmlroot functions
      - This matches other XML functions as well as the Postgres documentation,
        since these are closer to function argument names than regular keywords.
    - Fix deparse of ALTER TABLE a ALTER b SET STATISTICS DEFAULT
    - Fix the null pointer dereference when handling identity columns
  - Allow alternate definitions of NAMEDATALEN identifier limit
    - This allows building libpg_query with an override of the built-time limit of
      Postgres identifiers (typically 63 characters)
  - Bump Postgres to 17.7 and switch back to release tarballs
* Avoid conflicts with other cgo libraries utilizing xxhash by using XXH_NAMESPACE [#144](https://github.com/pganalyze/pg_query_go/pull/144)
* Remove redundant copy of result in ParseToProtobuf [#139](https://github.com/pganalyze/pg_query_go/pull/139)


## 6.1.0     2025-02-24

* Upgrade to libpg_query 17-6.1.0
  - Update to Postgres 17.4, and add recent patches scheduled for Postgres 17.5 (not yet released)
    - Notably, this pulls in support for macOS 15.4 which defines strchrnul
    in its standard library, fixing builds on up-to-date macOS versions.
  - Deparser improvements
    - Add parenthesis around AT LOCAL / AT TIMEZONE if needed
    - Correctness improvements related to expressions and function calls
* Allow vendoring pg_query_go with built-in "go mod vendor" command [(#131)](https://github.com/pganalyze/pg_query_go/pull/131)


## 6.0.0     2024-11-26

* Upgrade to libpg_query 17-6.0.0
  - Updates to the Postgres 17 parser
  - Deparser improvements:
    - Add support for deparsing `JSON_TABLE`, `JSON_QUERY`, `JSON_EXISTS`, `JSON_VALUE`
    - Add support for deparsing `JSON`, `JSON_SCALAR`, `JSON_SERIALIZE`
    - Add support for deparsing `COPY ... FORCE_NULL(*)`
    - Add support for deparsing `ALTER COLUMN ... SET EXPRESSION AS`
    - Add support for deparsing `SET STATISTICS DEFAULT`
    - Add support for deparsing `SET ACCESS METHOD DEFAULT`
    - Add support for deparsing `... AT LOCAL`
    - Add support for deparsing `merge_action()`
    - Add support for deparsing `MERGE ... RETURNING`
    - Add support for deparsing `NOT MATCHED [ BY TARGET ]`
* Add function `TestNormalizeUtility` to normalize only utility statements [(#116)](https://github.com/pganalyze/pg_query_go/pull/116)

## 5.1.0     2024-01-09

* Update to libpg_query 16-5.1.0
  - Add support for running on Windows
  - Add support for compiling on 32-bit systems


## 5.0.0     2023-12-23

* Upgrade to libpg_query 16-5.0.0
  - Updates to the Postgres 16 parser
  - Multiple deparser improvements
* Allow importing parse tree proto with cgo disabled
* Add parser.SplitWithScanner() and parser.SplitWithParser()


## 4.2.3     2023-08-04

* Update to libpg_query 15-4.2.3, including changes from 15-4.2.2
  - Fix builds when compiling with `glibc >=  2.38` [#203](https://github.com/pganalyze/libpg_query/pull/203)
  - Deparser: Add support for COALESCE and other expressions in LIMIT clause [#199](https://github.com/pganalyze/libpg_query/pull/199)
  - Deparser: Add support for multi-statement CREATE PROCEDURE definitions [#197](https://github.com/pganalyze/libpg_query/pull/197)
  - Deparser: Correctly quote identifier in ALTER TABLE ... ADD CONSTRAINT [x] [#196](https://github.com/pganalyze/libpg_query/pull/196)
  - Deparser: Add support for index fillfactor within CREATE TABLE, fix SHOW ALL [#193](https://github.com/pganalyze/libpg_query/pull/193)
  - Deparser: Move to dedicated file for easier inclusion in third-party projects [#192](https://github.com/pganalyze/libpg_query/pull/192)


## 4.2.1     2023-05-25

* Update to libpg_query 15-4.2.1 ([#85](https://github.com/pganalyze/pg_query_go/pull/85))
  - Deparser: Handle INTERVAL correctly when used in SET statements [#184](https://github.com/pganalyze/libpg_query/pull/184)
  - Deparser: Ensure index names are quoted as identifiers [#182](https://github.com/pganalyze/libpg_query/pull/182)
* Suppress -Wdeprecated-non-prototype warnings ([#83](https://github.com/pganalyze/pg_query_go/pull/83) and [#84](https://github.com/pganalyze/pg_query_go/pull/84))
* Return full structured error info instead of just error message ([#76](https://github.com/pganalyze/pg_query_go/pull/76))


## 4.2.0     2023-02-08

* Update to libpg_query 15-4.2.0
  - Update to PostgreSQL 15.1


## 2.2.0     2022-11-02

* Update to libpg_query 13-2.2.0 ([#64](https://github.com/pganalyze/pg_query_go/pull/64))
  - Fingerprinting version 3.1
    - Fixes issue with "SELECT DISTINCT" having the same fingerprint as "SELECT"
      (fingerprints for "SELECT DISTINCT" will change with this revision)
    - Group additional DDL statements together that otherwise generate a lot of
      unique fingerprints (ListenStmt, UnlistenStmt, NotifyStmt, CreateFunctionStmt,
      FunctionParameter and DoStmt)
  - Update to Postgres 13.8 patch release
  - Backport Xcode 14.1 build fix from upcoming 13.9 release
  - Normalize additional DDL statements


## 2.1.2      2022-06-28

* Update libpg_query to 13-2.1.2 ([#58](https://github.com/pganalyze/pg_query_go/pull/58))
  - Add support for analyzing PL/pgSQL code inside DO blocks [#142](https://github.com/pganalyze/libpg_query/pull/142)
  - Fix memory leak in pg_query_fingerprint error handling [#141](https://github.com/pganalyze/libpg_query/pull/141)
  - PL/pgSQL parser:
    - Add support for Assert [#135](https://github.com/pganalyze/libpg_query/pull/135)
    - Add support for SET, COMMIT, ROLLBACK and CALL [#130](https://github.com/pganalyze/libpg_query/pull/130)
  - Add support for parsing more operators that include a `?` character (special cased to support old pg_stat_statements query texts)
    - ltree extension [#136](https://github.com/pganalyze/libpg_query/pull/136)
    - promscale extension [#133](https://github.com/pganalyze/libpg_query/pull/133)
  - Deparser improvements
    - Prefix errors with "deparse", and remove some asserts [#131](https://github.com/pganalyze/libpg_query/pull/131)
    - Fix potential segfault when passing invalid protobuf (RawStmt without Stmt) [#128](https://github.com/pganalyze/libpg_query/pull/128)


## 2.1.0      2021-10-12

* Update libpg_query to 13-2.1.0 ([#53](https://github.com/pganalyze/pg_query_go/pull/53))
  - Normalize: add funcname error object
  - Normalize: Match GROUP BY against target list and re-use param refs
  - PL/pgSQL: Setup namespace items for parameters, support RECORD types
    - This significantly improves parsing for PL/pgSQL functions, to the
      extent that most functions should now parse successfully


## 2.0.5      2021-07-16

* Update libpg_query to 13-2.0.7 ([#49](https://github.com/pganalyze/pg_query_go/pull/49))
  - Normalize: Don't modify constants in TypeName typmods/arrayBounds fields
  - Don't fail builds on systems that have strchrnul support (FreeBSD)


## 2.0.4      2021-06-29

* Update libpg_query to 13-2.0.6 ([#47](https://github.com/pganalyze/pg_query_go/pull/47))
  - Normalize: Don't touch "ORDER BY 1" expressions, keep original text


## 2.0.3      2021-06-28

* Update libpg_query to 13-2.0.5 ([#45](https://github.com/pganalyze/pg_query_go/pull/45))
  - Update to Postgres 13.3 patch release
  - Normalize: Don't touch "GROUP BY 1" type statements, keep original text
  - Fingerprint: Cache list item hashes to fingerprint complex queries faster
  - Deparser: Emit the RangeVar catalogname if present
  - Fix crash in pg_scan function when encountering backslash escapes
* Add pg_query.Scan to access Postgres scanner ([#43](https://github.com/pganalyze/pg_query_go/pull/43))


## 2.0.2      2021-04-02

* Update to libpg_query 13-2.0.3
  - Normalize: Fix handling of two subsequent DefElem elements
* Parser CFLAGS: Build with -std=gnu99 to ensure CentOS compatibility


## 2.0.1      2021-03-30

* Update to libpg_query 13-2.0.2
  - Fix ARM builds: Avoid dependency on cpuid.h header
  - Simplify deparser of TableLikeClause
  - Fix asprintf warnings by ensuring _GNU_SOURCE is set early enough
* Add FingerprintToUInt64 method for callers that prefer to handle uint64s
  - This prevents a caller from having to do a hex string to uint64 conversion
    by simply returning the uint64 version of the fingerprint instead, which
    is already always provided by libpg_query.
* Add HashXXH3_64 helper method for generating XXH3 64-bit hash values
  - This can be useful when trying to fit other values into a data structure
    sized for the fingerprint, such as when encountering a parsing error
    during fingerprinting, and wanting to encode that fact uniquely into
    the fingerprint value.


## 2.0.0      2021-03-18

* Update libpg_query to 13-2.0.0
* Switch to use Protobuf generated nodes instead of custom logic
  - WARNING: This is breaking API change!
  - This is a breaking change in the API, but necessary in order to
    significantly improve the performance of parsing a query into Go structs,
    as well as allowing future bi-directional passing of parse trees between
    Go and C, such as for a future addition of a deparser.
* Rename pg_query.FastFingerprint to pg_query.Fingerprint
  - WARNING: This is breaking API change!
  - We no longer have direct support for running the fingerprint in the Go
    library, as its unnecessarily complex to support, and we can instead
    rely on the C fingerprinting method.
* Add support for deparsing parse trees back into a SQL statement
  - Call the new `pg_query.Deparse` method to get back the SQL text for
    a particular parse tree structure
  - This relies on the new deparser in libpg_query to transform a given parse
    tree back into a SQL statement. Note that this is currently considered
    experimental, and should not be used on unsanitized input, due to the risk
    of crashes in the C code with unexpected conditions.
* Update import path to `github.com/pganalyze/pg_query_go`


## 1.0.2      2021-02-18

* Update libpg_query to 10-1.0.5
  - This resolves memory leak problems, adds PPC architecture support,
    and refreshes the Postgres minor version to 10.16.


## 1.0.1      2020-11-07

* Update libpg_query to 10-1.0.3
  - This fixes ARM builds, and refreshes the Postgres patch version. This
    commit does not change the Postgres major version yet (still at PG 10).


## 1.0.0      2019-01-11

* Initial release with a tagged version
  - Note that 1.X releases will reflect Postgres 10 parser based releases from
    now on. When the Postgres 11 parser is released, that will be a new major
    version.
