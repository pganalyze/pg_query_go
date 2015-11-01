// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateOpFamilyStmt struct {
	Opfamilyname []Node  `json:"opfamilyname"` /* qualified name (list of Value strings) */
	Amname       *string `json:"amname"`       /* name of index AM opfamily is for */
}

func (node CreateOpFamilyStmt) MarshalJSON() ([]byte, error) {
	type CreateOpFamilyStmtMarshalAlias CreateOpFamilyStmt
	return json.Marshal(map[string]interface{}{
		"CREATEOPFAMILYSTMT": (*CreateOpFamilyStmtMarshalAlias)(&node),
	})
}

func (node *CreateOpFamilyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
