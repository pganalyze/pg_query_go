// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeSubselect struct {
	Lateral  bool   `json:"lateral"`  /* does it have LATERAL prefix? */
	Subquery Node   `json:"subquery"` /* the untransformed sub-select clause */
	Alias    *Alias `json:"alias"`    /* table alias & optional column aliases */
}

func (node RangeSubselect) MarshalJSON() ([]byte, error) {
	type RangeSubselectMarshalAlias RangeSubselect
	return json.Marshal(map[string]interface{}{
		"RANGESUBSELECT": (*RangeSubselectMarshalAlias)(&node),
	})
}

func (node *RangeSubselect) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
