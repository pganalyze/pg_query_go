// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterOpFamilyStmt struct {
	Opfamilyname []Node  `json:"opfamilyname"` /* qualified name (list of Value strings) */
	Amname       *string `json:"amname"`       /* name of index AM opfamily is for */
	IsDrop       bool    `json:"isDrop"`       /* ADD or DROP the items? */
	Items        []Node  `json:"items"`        /* List of CreateOpClassItem nodes */
}

func (node AlterOpFamilyStmt) MarshalJSON() ([]byte, error) {
	type AlterOpFamilyStmtMarshalAlias AlterOpFamilyStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROPFAMILYSTMT": (*AlterOpFamilyStmtMarshalAlias)(&node),
	})
}

func (node *AlterOpFamilyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
