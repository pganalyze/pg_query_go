// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterSeqStmt struct {
	Sequence  *RangeVar `json:"sequence"` /* the sequence to alter */
	Options   []Node    `json:"options"`
	MissingOk bool      `json:"missing_ok"` /* skip error if a role is missing? */
}

func (node AlterSeqStmt) MarshalJSON() ([]byte, error) {
	type AlterSeqStmtMarshalAlias AlterSeqStmt
	return json.Marshal(map[string]interface{}{
		"ALTERSEQSTMT": (*AlterSeqStmtMarshalAlias)(&node),
	})
}

func (node *AlterSeqStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
