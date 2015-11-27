// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeFunction struct {
	Lateral    bool   `json:"lateral"`     /* does it have LATERAL prefix? */
	Ordinality bool   `json:"ordinality"`  /* does it have WITH ORDINALITY suffix? */
	IsRowsfrom bool   `json:"is_rowsfrom"` /* is result of ROWS FROM() syntax? */
	Functions  []Node `json:"functions"`   /* per-function information, see above */
	Alias      *Alias `json:"alias"`       /* table alias & optional column aliases */
	Coldeflist []Node `json:"coldeflist"`  /* list of ColumnDef nodes to describe result
	 * of function returning RECORD */

}

func (node RangeFunction) MarshalJSON() ([]byte, error) {
	type RangeFunctionMarshalAlias RangeFunction
	return json.Marshal(map[string]interface{}{
		"RANGEFUNCTION": (*RangeFunctionMarshalAlias)(&node),
	})
}

func (node *RangeFunction) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["lateral"] != nil {
		err = json.Unmarshal(fields["lateral"], &node.Lateral)
		if err != nil {
			return
		}
	}

	if fields["ordinality"] != nil {
		err = json.Unmarshal(fields["ordinality"], &node.Ordinality)
		if err != nil {
			return
		}
	}

	if fields["is_rowsfrom"] != nil {
		err = json.Unmarshal(fields["is_rowsfrom"], &node.IsRowsfrom)
		if err != nil {
			return
		}
	}

	if fields["functions"] != nil {
		node.Functions, err = UnmarshalNodeArrayJSON(fields["functions"])
		if err != nil {
			return
		}
	}

	if fields["alias"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["alias"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Alias)
			node.Alias = &val
		}
	}

	if fields["coldeflist"] != nil {
		node.Coldeflist, err = UnmarshalNodeArrayJSON(fields["coldeflist"])
		if err != nil {
			return
		}
	}

	return
}
