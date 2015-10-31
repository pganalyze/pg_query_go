// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node RangeSubselect) Deparse() string {
	panic("Not Implemented")
}
