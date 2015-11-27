// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SortGroupClause struct {
	TleSortGroupRef Index `json:"tleSortGroupRef"` /* reference into targetlist */
	Eqop            Oid   `json:"eqop"`            /* the equality operator ('=' op) */
	Sortop          Oid   `json:"sortop"`          /* the ordering operator ('<' op), or 0 */
	NullsFirst      bool  `json:"nulls_first"`     /* do NULLs come before normal values? */
	Hashable        bool  `json:"hashable"`        /* can eqop be implemented by hashing? */
}

func (node SortGroupClause) MarshalJSON() ([]byte, error) {
	type SortGroupClauseMarshalAlias SortGroupClause
	return json.Marshal(map[string]interface{}{
		"SORTGROUPCLAUSE": (*SortGroupClauseMarshalAlias)(&node),
	})
}

func (node *SortGroupClause) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["tleSortGroupRef"] != nil {
		err = json.Unmarshal(fields["tleSortGroupRef"], &node.TleSortGroupRef)
		if err != nil {
			return
		}
	}

	if fields["eqop"] != nil {
		err = json.Unmarshal(fields["eqop"], &node.Eqop)
		if err != nil {
			return
		}
	}

	if fields["sortop"] != nil {
		err = json.Unmarshal(fields["sortop"], &node.Sortop)
		if err != nil {
			return
		}
	}

	if fields["nulls_first"] != nil {
		err = json.Unmarshal(fields["nulls_first"], &node.NullsFirst)
		if err != nil {
			return
		}
	}

	if fields["hashable"] != nil {
		err = json.Unmarshal(fields["hashable"], &node.Hashable)
		if err != nil {
			return
		}
	}

	return
}
