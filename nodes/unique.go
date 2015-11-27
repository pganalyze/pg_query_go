// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type Unique struct {
	Plan          Plan        `json:"plan"`
	NumCols       int         `json:"numCols"`       /* number of columns to check for uniqueness */
	UniqColIdx    *AttrNumber `json:"uniqColIdx"`    /* their indexes in the target list */
	UniqOperators *Oid        `json:"uniqOperators"` /* equality operators to compare with */
}

func (node Unique) MarshalJSON() ([]byte, error) {
	type UniqueMarshalAlias Unique
	return json.Marshal(map[string]interface{}{
		"UNIQUE": (*UniqueMarshalAlias)(&node),
	})
}

func (node *Unique) UnmarshalJSON(input []byte) (err error) {
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

	if fields["uniqColIdx"] != nil {
		err = json.Unmarshal(fields["uniqColIdx"], &node.UniqColIdx)
		if err != nil {
			return
		}
	}

	if fields["uniqOperators"] != nil {
		err = json.Unmarshal(fields["uniqOperators"], &node.UniqOperators)
		if err != nil {
			return
		}
	}

	return
}
