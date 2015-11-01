// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateSeqStmt struct {
	Sequence *RangeVar `json:"sequence"` /* the sequence to create */
	Options  []Node    `json:"options"`
	OwnerId  Oid       `json:"ownerId"` /* ID of owner, or InvalidOid for default */
}

func (node CreateSeqStmt) MarshalJSON() ([]byte, error) {
	type CreateSeqStmtMarshalAlias CreateSeqStmt
	return json.Marshal(map[string]interface{}{
		"CREATESEQSTMT": (*CreateSeqStmtMarshalAlias)(&node),
	})
}

func (node *CreateSeqStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
