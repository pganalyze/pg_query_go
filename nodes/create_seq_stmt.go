// Auto-generated - DO NOT EDIT

package pg_query

type CreateSeqStmt struct {
	Sequence *RangeVar `json:"sequence"` /* the sequence to create */
	Options  []Node    `json:"options"`
	OwnerId  Oid       `json:"ownerId"` /* ID of owner, or InvalidOid for default */
}
