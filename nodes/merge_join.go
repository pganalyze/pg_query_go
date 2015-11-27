// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type MergeJoin struct {
	Join         Join   `json:"join"`
	Mergeclauses []Node `json:"mergeclauses"` /* mergeclauses as expression trees */

	/* these are arrays, but have the same length as the mergeclauses list: */
	MergeFamilies   *Oid  `json:"mergeFamilies"`   /* per-clause OIDs of btree opfamilies */
	MergeCollations *Oid  `json:"mergeCollations"` /* per-clause OIDs of collations */
	MergeStrategies *int  `json:"mergeStrategies"` /* per-clause ordering (ASC or DESC) */
	MergeNullsFirst *bool `json:"mergeNullsFirst"` /* per-clause nulls ordering */
}

func (node MergeJoin) MarshalJSON() ([]byte, error) {
	type MergeJoinMarshalAlias MergeJoin
	return json.Marshal(map[string]interface{}{
		"MERGEJOIN": (*MergeJoinMarshalAlias)(&node),
	})
}

func (node *MergeJoin) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["join"] != nil {
		err = json.Unmarshal(fields["join"], &node.Join)
		if err != nil {
			return
		}
	}

	if fields["mergeclauses"] != nil {
		node.Mergeclauses, err = UnmarshalNodeArrayJSON(fields["mergeclauses"])
		if err != nil {
			return
		}
	}

	if fields["mergeFamilies"] != nil {
		err = json.Unmarshal(fields["mergeFamilies"], &node.MergeFamilies)
		if err != nil {
			return
		}
	}

	if fields["mergeCollations"] != nil {
		err = json.Unmarshal(fields["mergeCollations"], &node.MergeCollations)
		if err != nil {
			return
		}
	}

	if fields["mergeStrategies"] != nil {
		err = json.Unmarshal(fields["mergeStrategies"], &node.MergeStrategies)
		if err != nil {
			return
		}
	}

	if fields["mergeNullsFirst"] != nil {
		err = json.Unmarshal(fields["mergeNullsFirst"], &node.MergeNullsFirst)
		if err != nil {
			return
		}
	}

	return
}
