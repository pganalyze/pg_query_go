// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RuleStmt) MarshalJSON() ([]byte, error) {
	type RuleStmtMarshalAlias RuleStmt
	return json.Marshal(map[string]interface{}{
		"RULESTMT": (*RuleStmtMarshalAlias)(&node),
	})
}

func (node *RuleStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RuleStmt) Deparse() string {
	panic("Not Implemented")
}
