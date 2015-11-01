// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WindowDef struct {
	Name            *string `json:"name"`            /* window's own name */
	Refname         *string `json:"refname"`         /* referenced window name, if any */
	PartitionClause []Node  `json:"partitionClause"` /* PARTITION BY expression list */
	OrderClause     []Node  `json:"orderClause"`     /* ORDER BY (list of SortBy) */
	FrameOptions    int     `json:"frameOptions"`    /* frame_clause options, see below */
	StartOffset     Node    `json:"startOffset"`     /* expression for starting bound, if any */
	EndOffset       Node    `json:"endOffset"`       /* expression for ending bound, if any */
	Location        int     `json:"location"`        /* parse location, or -1 if none/unknown */
}

func (node WindowDef) MarshalJSON() ([]byte, error) {
	type WindowDefMarshalAlias WindowDef
	return json.Marshal(map[string]interface{}{
		"WINDOWDEF": (*WindowDefMarshalAlias)(&node),
	})
}

func (node *WindowDef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
