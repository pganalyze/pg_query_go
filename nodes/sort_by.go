// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SortBy struct {
	Node        Node        `json:"node"`         /* expression to sort on */
	SortbyDir   SortByDir   `json:"sortby_dir"`   /* ASC/DESC/USING/default */
	SortbyNulls SortByNulls `json:"sortby_nulls"` /* NULLS FIRST/LAST */
	UseOp       []Node      `json:"useOp"`        /* name of op to use, if SORTBY_USING */
	Location    int         `json:"location"`     /* operator location, or -1 if none/unknown */
}

func (node SortBy) MarshalJSON() ([]byte, error) {
	type SortByMarshalAlias SortBy
	return json.Marshal(map[string]interface{}{
		"SORTBY": (*SortByMarshalAlias)(&node),
	})
}

func (node *SortBy) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
