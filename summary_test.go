//go:build cgo
// +build cgo

package pg_query_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

var summaryTests = []struct {
	input         string
	truncateLimit int
	expected      *pg_query.SummaryResult
}{
	// Basic SELECT with filter column
	{
		input:         "SELECT * FROM users WHERE id = 1",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_Select}},
			FilterColumns:  []*pg_query.SummaryResult_FilterColumn{{Column: "id"}},
			StatementTypes: []string{"SelectStmt"},
		},
	},
	// Query truncation
	{
		input:         "SELECT id, name, email FROM users WHERE id = 1",
		truncateLimit: 30,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_Select}},
			FilterColumns:  []*pg_query.SummaryResult_FilterColumn{{Column: "id"}},
			StatementTypes: []string{"SelectStmt"},
			TruncatedQuery: "SELECT ... FROM users WHERE...",
		},
	},
	// JOIN with table aliases
	{
		input:         "SELECT * FROM users u JOIN orders o ON u.id = o.user_id",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables: []*pg_query.SummaryResult_Table{
				{Name: "users", TableName: "users", Context: pg_query.SummaryResult_Select},
				{Name: "orders", TableName: "orders", Context: pg_query.SummaryResult_Select},
			},
			Aliases:        map[string]string{"u": "users", "o": "orders"},
			StatementTypes: []string{"SelectStmt"},
		},
	},
	// Aggregate functions (count, max)
	{
		input:         "SELECT count(*), max(id) FROM users",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables: []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_Select}},
			Functions: []*pg_query.SummaryResult_Function{
				{Name: "count", FunctionName: "count", Context: pg_query.SummaryResult_Call},
				{Name: "max", FunctionName: "max", Context: pg_query.SummaryResult_Call},
			},
			StatementTypes: []string{"SelectStmt"},
		},
	},
	// Common Table Expression (CTE)
	{
		input:         "WITH active_users AS (SELECT * FROM users WHERE active = true) SELECT * FROM active_users",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_Select}},
			CteNames:       []string{"active_users"},
			FilterColumns:  []*pg_query.SummaryResult_FilterColumn{{Column: "active"}},
			StatementTypes: []string{"SelectStmt"},
		},
	},
	// Schema-qualified table name
	{
		input:         "SELECT * FROM public.users",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "public.users", SchemaName: "public", TableName: "users", Context: pg_query.SummaryResult_Select}},
			StatementTypes: []string{"SelectStmt"},
		},
	},
	// INSERT statement
	{
		input:         "INSERT INTO users (name) VALUES ('test')",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_DML}},
			StatementTypes: []string{"InsertStmt", "SelectStmt"},
		},
	},
	// UPDATE statement
	{
		input:         "UPDATE users SET name = 'test' WHERE id = 1",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_DML}},
			StatementTypes: []string{"UpdateStmt"},
		},
	},
	// DELETE statement
	{
		input:         "DELETE FROM users WHERE id = 1",
		truncateLimit: -1,
		expected: &pg_query.SummaryResult{
			Tables:         []*pg_query.SummaryResult_Table{{Name: "users", TableName: "users", Context: pg_query.SummaryResult_DML}},
			StatementTypes: []string{"DeleteStmt"},
		},
	},
}

func TestSummary(t *testing.T) {
	for _, tt := range summaryTests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := pg_query.Summary(tt.input, tt.truncateLimit)
			if err != nil {
				t.Fatalf("Summary() error: %s", err)
			}

			if !proto.Equal(result, tt.expected) {
				t.Errorf("Summary() mismatch:\n  got:      %v\n  expected: %v", result, tt.expected)
			}
		})
	}
}

func TestSummaryError(t *testing.T) {
	_, err := pg_query.Summary("SELECT * FROM", -1)
	if err == nil {
		t.Error("expected error for invalid SQL")
	}
}
