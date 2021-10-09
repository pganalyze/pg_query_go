package pg_query_test

import (
	"testing"

	_ "github.com/pganalyze/pg_query_go/v2"
	pg_query "github.com/pganalyze/pg_query_go/v2"
)

func Test(t *testing.T) {
	result, err := pg_query.SplitWithScanner("select 1; select 2; select 3;")
	if err != nil {
		t.Error(err)
	} else {
		expected := []string{"select 1", "select 2", "select 3"}
		for i, res := range result {
			if expected[i] != res {
				t.Errorf("Unexpected split: %v", result)
			}
		}
	}
}
