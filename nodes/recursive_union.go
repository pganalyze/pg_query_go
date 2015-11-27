// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RecursiveUnion struct {
	Plan    Plan `json:"plan"`
	WtParam int  `json:"wtParam"` /* ID of Param representing work table */

	/* Remaining fields are zero/null in UNION ALL case */
	NumCols int `json:"numCols"` /* number of columns to check for
	 * duplicate-ness */

	DupColIdx    *AttrNumber `json:"dupColIdx"`    /* their indexes in the target list */
	DupOperators *Oid        `json:"dupOperators"` /* equality operators to compare with */
	NumGroups    int64       `json:"numGroups"`    /* estimated number of groups in input */
}

func (node RecursiveUnion) MarshalJSON() ([]byte, error) {
	type RecursiveUnionMarshalAlias RecursiveUnion
	return json.Marshal(map[string]interface{}{
		"RECURSIVEUNION": (*RecursiveUnionMarshalAlias)(&node),
	})
}

func (node *RecursiveUnion) UnmarshalJSON(input []byte) (err error) {
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

	if fields["wtParam"] != nil {
		err = json.Unmarshal(fields["wtParam"], &node.WtParam)
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

	if fields["dupColIdx"] != nil {
		err = json.Unmarshal(fields["dupColIdx"], &node.DupColIdx)
		if err != nil {
			return
		}
	}

	if fields["dupOperators"] != nil {
		err = json.Unmarshal(fields["dupOperators"], &node.DupOperators)
		if err != nil {
			return
		}
	}

	if fields["numGroups"] != nil {
		err = json.Unmarshal(fields["numGroups"], &node.NumGroups)
		if err != nil {
			return
		}
	}

	return
}
