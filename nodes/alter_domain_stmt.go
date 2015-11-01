// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterDomainStmt struct {
	Subtype byte `json:"subtype"` /*------------
	 *	T = alter column default
	 *	N = alter column drop not null
	 *	O = alter column set not null
	 *	C = add constraint
	 *	X = drop constraint
	 *------------
	 */
	TypeName  []Node       `json:"typeName"`   /* domain to work on */
	Name      *string      `json:"name"`       /* column or constraint name to act on */
	Def       Node         `json:"def"`        /* definition of default or constraint */
	Behavior  DropBehavior `json:"behavior"`   /* RESTRICT or CASCADE for DROP cases */
	MissingOk bool         `json:"missing_ok"` /* skip error if missing? */
}

func (node AlterDomainStmt) MarshalJSON() ([]byte, error) {
	type AlterDomainStmtMarshalAlias AlterDomainStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDOMAINSTMT": (*AlterDomainStmtMarshalAlias)(&node),
	})
}

func (node *AlterDomainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
