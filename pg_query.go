//go:build cgo
// +build cgo

package pg_query

import (
	"google.golang.org/protobuf/proto"

	"github.com/pganalyze/pg_query_go/v6/parser"
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

// ParseToJSON - Parses the given SQL statement into a parse tree (JSON format)
func ParseToJSON(input string) (result string, err error) {
	return parser.ParseToJSON(input)
}

// Parse the given SQL statement into a parse tree (Go struct format)
func Parse(input string) (tree *ParseResult, err error) {
	protobufTree, err := parser.ParseToProtobuf(input)
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

// FingerprintOption is a bitmask of flags controlling how fingerprints are
// calculated, mirroring the PgQueryFingerprintOption enum in pg_query.h
type FingerprintOption = parser.FingerprintOption

const (
	// FingerprintDefault fingerprints relation references following the
	// Postgres 18+ query ID behavior: in SELECT/DML statements the alias name
	// is fingerprinted, the relation name is ignored when an alias is present,
	// and schema names are ignored. Sequences of two or more digits in
	// relation names are ignored, so that queries on date/number-suffixed
	// tables (e.g. partitions like "orders_2024_01") get the same fingerprint.
	FingerprintDefault = parser.FingerprintDefault

	// FingerprintRangeVarIgnoreAliases always fingerprints relation names and ignores aliases
	FingerprintRangeVarIgnoreAliases = parser.FingerprintRangeVarIgnoreAliases

	// FingerprintRangeVarIncludeSchema also fingerprints schema names in SELECT/DML
	// statements (they are always fingerprinted in utility statements)
	FingerprintRangeVarIncludeSchema = parser.FingerprintRangeVarIncludeSchema

	// FingerprintRangeVarPG17Compat is a convenience combination that matches how
	// Postgres 17 and earlier calculate query IDs, and how libpg_query 17 and
	// earlier calculated fingerprints
	FingerprintRangeVarPG17Compat = parser.FingerprintRangeVarPG17Compat

	// FingerprintRelnameFull fingerprints the full relation name, including digit sequences
	FingerprintRelnameFull = parser.FingerprintRelnameFull
)

// Fingerprint - Fingerprint the passed SQL statement to a hex string
func Fingerprint(input string) (result string, err error) {
	return parser.FingerprintToHexStr(input)
}

// FingerprintWithOpts - Fingerprint the passed SQL statement to a hex string,
// with options controlling how the fingerprint is calculated
func FingerprintWithOpts(input string, opts FingerprintOption) (result string, err error) {
	return parser.FingerprintToHexStrWithOpts(input, opts)
}

// FingerprintToUInt64 - Fingerprint the passed SQL statement to a uint64
func FingerprintToUInt64(input string) (result uint64, err error) {
	return parser.FingerprintToUInt64(input)
}

// FingerprintToUInt64WithOpts - Fingerprint the passed SQL statement to a uint64,
// with options controlling how the fingerprint is calculated
func FingerprintToUInt64WithOpts(input string, opts FingerprintOption) (result uint64, err error) {
	return parser.FingerprintToUInt64WithOpts(input, opts)
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
func Summary(input string, truncateLimit int) (result *SummaryResult, err error) {
	protobufSummary, err := parser.SummaryToProtobuf(input, truncateLimit)
	if err != nil {
		return
	}
	result = &SummaryResult{}
	err = proto.Unmarshal(protobufSummary, result)
	return
}
