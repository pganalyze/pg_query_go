// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WindowAgg struct {
	Plan          Plan        `json:"plan"`
	Winref        Index       `json:"winref"`        /* ID referenced by window functions */
	PartNumCols   int         `json:"partNumCols"`   /* number of columns in partition clause */
	PartColIdx    *AttrNumber `json:"partColIdx"`    /* their indexes in the target list */
	PartOperators *Oid        `json:"partOperators"` /* equality operators for partition columns */
	OrdNumCols    int         `json:"ordNumCols"`    /* number of columns in ordering clause */
	OrdColIdx     *AttrNumber `json:"ordColIdx"`     /* their indexes in the target list */
	OrdOperators  *Oid        `json:"ordOperators"`  /* equality operators for ordering columns */
	FrameOptions  int         `json:"frameOptions"`  /* frame_clause options, see WindowDef */
	StartOffset   Node        `json:"startOffset"`   /* expression for starting bound, if any */
	EndOffset     Node        `json:"endOffset"`     /* expression for ending bound, if any */
}

func (node WindowAgg) MarshalJSON() ([]byte, error) {
	type WindowAggMarshalAlias WindowAgg
	return json.Marshal(map[string]interface{}{
		"WINDOWAGG": (*WindowAggMarshalAlias)(&node),
	})
}

func (node *WindowAgg) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
