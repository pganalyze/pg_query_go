// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SortGroupClause) MarshalJSON() ([]byte, error) {
	type SortGroupClauseMarshalAlias SortGroupClause
	return json.Marshal(map[string]interface{}{
		"SORTGROUPCLAUSE": (*SortGroupClauseMarshalAlias)(&node),
	})
}

func (node *SortGroupClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SortGroupClause) Deparse() string {
	panic("Not Implemented")
}
