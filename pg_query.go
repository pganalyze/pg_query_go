//go:build cgo
// +build cgo

package pg_query

import (
	"google.golang.org/protobuf/proto"

	"github.com/pganalyze/pg_query_go/v18/parser"
)

func Scan(input string) (result *ScanResult, err error) {
	protobufScan, err := parser.ScanToProtobuf(input)
	if err != nil {
		return
	}
	result = &ScanResult{}
	err = proto.Unmarshal(protobufScan, result)
	return
}

// ParserOptions control how a query string is parsed, mirroring the parser
// options of the C pg_query_parse_opts / pg_query_fingerprint_opts API.
//
// The parse mode (which grammar to use) occupies the lower 4 bits, the
// GUC-style flags occupy the bits above them, and both can be combined with
// a bitwise OR (e.g. ParseTypeName | DisableBackslashQuote).
//
// The values match the PG_QUERY_PARSE_* and PG_QUERY_DISABLE_* definitions
// in pg_query.h.
//
// Pass 0 (ParseDefault) when no specific options are needed.
type ParserOptions int

const (
	// ParseDefault parses using the standard top-level SQL grammar (the default)
	ParseDefault ParserOptions = 0
	// ParseTypeName parses a type name (e.g. "integer", "character varying(32)")
	ParseTypeName ParserOptions = 1
	// ParsePlPgSqlExpr parses a PL/pgSQL expression
	ParsePlPgSqlExpr ParserOptions = 2
	// ParsePlPgSqlAssign1 parses a PL/pgSQL assignment (target 1)
	ParsePlPgSqlAssign1 ParserOptions = 3
	// ParsePlPgSqlAssign2 parses a PL/pgSQL assignment (target 2)
	ParsePlPgSqlAssign2 ParserOptions = 4
	// ParsePlPgSqlAssign3 parses a PL/pgSQL assignment (target 3)
	ParsePlPgSqlAssign3 ParserOptions = 5

	// DisableBackslashQuote sets backslash_quote = off (default is safe_encoding, which is effectively on)
	DisableBackslashQuote ParserOptions = 16
	// DisableStandardConformingStrings sets standard_conforming_strings = off (default is on)
	DisableStandardConformingStrings ParserOptions = 32
	// DisableEscapeStringWarning sets escape_string_warning = off (default is on)
	DisableEscapeStringWarning ParserOptions = 64
)

// ParseToJSON - Parses the given SQL statement into a parse tree (JSON format)
//
// The parserOptions control how the statement is parsed, pass 0 for the
// standard top-level SQL grammar.
func ParseToJSON(input string, parserOptions ParserOptions) (result string, err error) {
	return parser.ParseToJSON(input, int(parserOptions))
}

// Parse the given SQL statement into a parse tree (Go struct format)
//
// The parserOptions control how the statement is parsed, pass 0 for the
// standard top-level SQL grammar.
func Parse(input string, parserOptions ParserOptions) (tree *ParseResult, err error) {
	protobufTree, err := parser.ParseToProtobuf(input, int(parserOptions))
	if err != nil {
		return
	}

	tree = &ParseResult{}
	err = proto.Unmarshal(protobufTree, tree)
	return
}

// Deparses a given Go parse tree into a SQL statement
func Deparse(tree *ParseResult) (output string, err error) {
	protobufTree, err := proto.Marshal(tree)
	if err != nil {
		return
	}

	output, err = parser.DeparseFromProtobuf(protobufTree)
	return
}

// ParsePlPgSqlToJSON - Parses the given PL/pgSQL function statement into a parse tree (JSON format)
func ParsePlPgSqlToJSON(input string) (result string, err error) {
	return parser.ParsePlPgSqlToJSON(input)
}

// Normalize the passed SQL statement to replace constant values with $n parameter references
func Normalize(input string) (result string, err error) {
	return parser.Normalize(input)
}

// Normalize the passed utility statement to replace constant values with $n parameter references
func NormalizeUtility(input string) (result string, err error) {
	return parser.NormalizeUtility(input)
}

// FingerprintOptions is a bitmask of flags controlling how fingerprints are
// calculated, mirroring the fingerprint options of the C
// pg_query_fingerprint_opts API. Flags can be combined with a bitwise OR
// (e.g. FingerprintRangeVarPG17Compat | FingerprintFullRelname).
//
// The values match the PG_QUERY_FINGERPRINT_* definitions in pg_query.h.
//
// Pass 0 (FingerprintDefault) when no specific options are needed.
type FingerprintOptions int

const (
	// FingerprintDefault fingerprints relation references following the
	// Postgres 18+ query ID behavior: in SELECT/DML statements the alias name
	// is fingerprinted, the relation name is ignored when an alias is present,
	// and schema names are ignored. Sequences of two or more digits in
	// relation names are ignored, so that queries on date/number-suffixed
	// tables (e.g. partitions like "orders_2024_01") get the same fingerprint.
	FingerprintDefault FingerprintOptions = 0

	// FingerprintRangeVarIgnoreAliases always fingerprints relation names and ignores aliases
	FingerprintRangeVarIgnoreAliases FingerprintOptions = 1 << 0

	// FingerprintRangeVarIncludeSchema also fingerprints schema names in SELECT/DML
	// statements (they are always fingerprinted in utility statements)
	FingerprintRangeVarIncludeSchema FingerprintOptions = 1 << 1

	// FingerprintRangeVarPG17Compat is a convenience combination that matches how
	// Postgres 17 and earlier calculate query IDs, and how libpg_query 17 and
	// earlier calculated fingerprints
	FingerprintRangeVarPG17Compat FingerprintOptions = FingerprintRangeVarIgnoreAliases | FingerprintRangeVarIncludeSchema

	// FingerprintFullRelname fingerprints the full relation name, including digit sequences
	FingerprintFullRelname FingerprintOptions = 1 << 4
)

// Fingerprint - Fingerprint the passed SQL statement to a hex string
//
// The parserOptions control how the statement is parsed, and the
// fingerprintOptions control how the fingerprint is calculated. Pass 0 for
// either when no specific options are needed.
//
// Note that as of Postgres 18 the default fingerprint ignores schema names and
// uses table aliases in place of relation names in SELECT/DML statements. Pass
// FingerprintRangeVarPG17Compat to get fingerprints matching earlier releases.
func Fingerprint(input string, parserOptions ParserOptions, fingerprintOptions FingerprintOptions) (result string, err error) {
	return parser.FingerprintToHexStr(input, int(parserOptions), int(fingerprintOptions))
}

// FingerprintToUInt64 - Fingerprint the passed SQL statement to a uint64
//
// See Fingerprint for details on the options.
func FingerprintToUInt64(input string, parserOptions ParserOptions, fingerprintOptions FingerprintOptions) (result uint64, err error) {
	return parser.FingerprintToUInt64(input, int(parserOptions), int(fingerprintOptions))
}

// HashXXH3_64 - Helper method to run XXH3 hash function (64-bit variant) on the given bytes, with the specified seed
func HashXXH3_64(input []byte, seed uint64) (result uint64) {
	return parser.HashXXH3_64(input, seed)
}

func SplitWithScanner(input string, trimSpace bool) (result []string, err error) {
	return parser.SplitWithScanner(input, trimSpace)
}

func SplitWithParser(input string, trimSpace bool) (result []string, err error) {
	return parser.SplitWithParser(input, trimSpace)
}

// IsUtilityStmt - Determines whether each statement in the query is a utility statement
//
// Returns a slice of booleans, one for each statement in the query.
// true = utility statement / DDL, false = SELECT / INSERT / UPDATE / DELETE / MERGE
func IsUtilityStmt(input string) (result []bool, err error) {
	return parser.IsUtilityStmt(input)
}

// Summary - Extracts summary information from SQL statement
//
// Optionally, you can pass a positive numbered truncateLimit to return a
// "smart" truncated version of the input statement that is at most limit length.
// Pass -1 to not truncate.
//
// The parserOptions control how the statement is parsed, pass 0 for the
// standard top-level SQL grammar.
func Summary(input string, parserOptions ParserOptions, truncateLimit int) (result *SummaryResult, err error) {
	protobufSummary, err := parser.SummaryToProtobuf(input, int(parserOptions), truncateLimit)
	if err != nil {
		return
	}
	result = &SummaryResult{}
	err = proto.Unmarshal(protobufSummary, result)
	return
}
