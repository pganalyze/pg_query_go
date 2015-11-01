// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type IntoClause struct {
	Rel            *RangeVar      `json:"rel"`            /* target relation name */
	ColNames       []Node         `json:"colNames"`       /* column names to assign, or NIL */
	Options        []Node         `json:"options"`        /* options from WITH clause */
	OnCommit       OnCommitAction `json:"onCommit"`       /* what do we do at COMMIT? */
	TableSpaceName *string        `json:"tableSpaceName"` /* table space to use, or NULL */
	ViewQuery      Node           `json:"viewQuery"`      /* materialized view's SELECT query */
	SkipData       bool           `json:"skipData"`       /* true for WITH NO DATA */
}

func (node IntoClause) MarshalJSON() ([]byte, error) {
	type IntoClauseMarshalAlias IntoClause
	return json.Marshal(map[string]interface{}{
		"INTOCLAUSE": (*IntoClauseMarshalAlias)(&node),
	})
}

func (node *IntoClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
