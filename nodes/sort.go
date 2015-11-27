// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type Sort struct {
	Plan          Plan        `json:"plan"`
	NumCols       int         `json:"numCols"`       /* number of sort-key columns */
	SortColIdx    *AttrNumber `json:"sortColIdx"`    /* their indexes in the target list */
	SortOperators *Oid        `json:"sortOperators"` /* OIDs of operators to sort them by */
	Collations    *Oid        `json:"collations"`    /* OIDs of collations */
	NullsFirst    *bool       `json:"nullsFirst"`    /* NULLS FIRST/LAST directions */
}

func (node Sort) MarshalJSON() ([]byte, error) {
	type SortMarshalAlias Sort
	return json.Marshal(map[string]interface{}{
		"SORT": (*SortMarshalAlias)(&node),
	})
}

func (node *Sort) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["numCols"] != nil {
		err = json.Unmarshal(fields["numCols"], &node.NumCols)
		if err != nil {
			return
		}
	}

	if fields["sortColIdx"] != nil {
		err = json.Unmarshal(fields["sortColIdx"], &node.SortColIdx)
		if err != nil {
			return
		}
	}

	if fields["sortOperators"] != nil {
		err = json.Unmarshal(fields["sortOperators"], &node.SortOperators)
		if err != nil {
			return
		}
	}

	if fields["collations"] != nil {
		err = json.Unmarshal(fields["collations"], &node.Collations)
		if err != nil {
			return
		}
	}

	if fields["nullsFirst"] != nil {
		err = json.Unmarshal(fields["nullsFirst"], &node.NullsFirst)
		if err != nil {
			return
		}
	}

	return
}
