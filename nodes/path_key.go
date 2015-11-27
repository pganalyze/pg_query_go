// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PathKey struct {
	PkEclass     *EquivalenceClass `json:"pk_eclass"`      /* the value that is ordered */
	PkOpfamily   Oid               `json:"pk_opfamily"`    /* btree opfamily defining the ordering */
	PkStrategy   int               `json:"pk_strategy"`    /* sort direction (ASC or DESC) */
	PkNullsFirst bool              `json:"pk_nulls_first"` /* do NULLs come before normal values? */
}

func (node PathKey) MarshalJSON() ([]byte, error) {
	type PathKeyMarshalAlias PathKey
	return json.Marshal(map[string]interface{}{
		"PATHKEY": (*PathKeyMarshalAlias)(&node),
	})
}

func (node *PathKey) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["pk_eclass"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["pk_eclass"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.PkEclass = &val
		}
	}

	if fields["pk_opfamily"] != nil {
		err = json.Unmarshal(fields["pk_opfamily"], &node.PkOpfamily)
		if err != nil {
			return
		}
	}

	if fields["pk_strategy"] != nil {
		err = json.Unmarshal(fields["pk_strategy"], &node.PkStrategy)
		if err != nil {
			return
		}
	}

	if fields["pk_nulls_first"] != nil {
		err = json.Unmarshal(fields["pk_nulls_first"], &node.PkNullsFirst)
		if err != nil {
			return
		}
	}

	return
}
