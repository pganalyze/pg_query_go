// Auto-generated - DO NOT EDIT

package pg_query

type RuleStmt struct {
	Relation    *RangeVar `json:"relation"`    /* relation the rule is for */
	Rulename    *string   `json:"rulename"`    /* name of the rule */
	WhereClause Node      `json:"whereClause"` /* qualifications */
	Event       CmdType   `json:"event"`       /* SELECT, INSERT, etc */
	Instead     bool      `json:"instead"`     /* is a 'do instead'? */
	Actions     []Node    `json:"actions"`     /* the action statements */
	Replace     bool      `json:"replace"`     /* OR REPLACE */
}
