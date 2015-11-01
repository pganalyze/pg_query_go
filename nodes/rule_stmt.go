// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RuleStmt struct {
	Relation    *RangeVar `json:"relation"`    /* relation the rule is for */
	Rulename    *string   `json:"rulename"`    /* name of the rule */
	WhereClause Node      `json:"whereClause"` /* qualifications */
	Event       CmdType   `json:"event"`       /* SELECT, INSERT, etc */
	Instead     bool      `json:"instead"`     /* is a 'do instead'? */
	Actions     []Node    `json:"actions"`     /* the action statements */
	Replace     bool      `json:"replace"`     /* OR REPLACE */
}

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
