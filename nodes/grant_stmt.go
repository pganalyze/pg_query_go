// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type GrantStmt struct {
	IsGrant  bool            `json:"is_grant"` /* true = GRANT, false = REVOKE */
	Targtype GrantTargetType `json:"targtype"` /* type of the grant target */
	Objtype  GrantObjectType `json:"objtype"`  /* kind of object being operated on */
	Objects  []Node          `json:"objects"`  /* list of RangeVar nodes, FuncWithArgs nodes,
	 * or plain names (as Value strings) */
	Privileges []Node `json:"privileges"` /* list of AccessPriv nodes */
	/* privileges == NIL denotes ALL PRIVILEGES */
	Grantees    []Node       `json:"grantees"`     /* list of PrivGrantee nodes */
	GrantOption bool         `json:"grant_option"` /* grant or revoke grant option */
	Behavior    DropBehavior `json:"behavior"`     /* drop behavior (for REVOKE) */
}

func (node GrantStmt) MarshalJSON() ([]byte, error) {
	type GrantStmtMarshalAlias GrantStmt
	return json.Marshal(map[string]interface{}{
		"GRANTSTMT": (*GrantStmtMarshalAlias)(&node),
	})
}

func (node *GrantStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
