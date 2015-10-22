package pg_query

type IntoClause struct {
	Rel            *RangeVar      `json:"rel"`            /* target relation name */
	ColNames       []Node         `json:"colNames"`       /* column names to assign, or NIL */
	Options        []Node         `json:"options"`        /* options from WITH clause */
	OnCommit       OnCommitAction `json:"onCommit"`       /* what do we do at COMMIT? */
	TableSpaceName *string        `json:"tableSpaceName"` /* table space to use, or NULL */
	ViewQuery      *Node          `json:"viewQuery"`      /* materialized view's SELECT query */
	SkipData       bool           `json:"skipData"`       /* true for WITH NO DATA */
}

func (intoClause IntoClause) Deparse() string {
	panic("Not Implemented")
}
