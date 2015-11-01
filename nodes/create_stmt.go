// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateStmt struct {
	Relation     *RangeVar `json:"relation"`     /* relation to create */
	TableElts    []Node    `json:"tableElts"`    /* column definitions (list of ColumnDef) */
	InhRelations []Node    `json:"inhRelations"` /* relations to inherit from (list of
	 * inhRelation) */
	OfTypename     *TypeName      `json:"ofTypename"`     /* OF typename */
	Constraints    []Node         `json:"constraints"`    /* constraints (list of Constraint nodes) */
	Options        []Node         `json:"options"`        /* options from WITH clause */
	Oncommit       OnCommitAction `json:"oncommit"`       /* what do we do at COMMIT? */
	Tablespacename *string        `json:"tablespacename"` /* table space to use, or NULL */
	IfNotExists    bool           `json:"if_not_exists"`  /* just do nothing if it already exists? */
}

func (node CreateStmt) MarshalJSON() ([]byte, error) {
	type CreateStmtMarshalAlias CreateStmt
	return json.Marshal(map[string]interface{}{
		"CREATESTMT": (*CreateStmtMarshalAlias)(&node),
	})
}

func (node *CreateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
