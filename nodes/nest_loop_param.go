// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type NestLoopParam struct {
	Paramno  int  `json:"paramno"`  /* number of the PARAM_EXEC Param to set */
	Paramval *Var `json:"paramval"` /* outer-relation Var to assign to Param */
}

func (node NestLoopParam) MarshalJSON() ([]byte, error) {
	type NestLoopParamMarshalAlias NestLoopParam
	return json.Marshal(map[string]interface{}{
		"NESTLOOPPARAM": (*NestLoopParamMarshalAlias)(&node),
	})
}

func (node *NestLoopParam) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["paramno"] != nil {
		err = json.Unmarshal(fields["paramno"], &node.Paramno)
		if err != nil {
			return
		}
	}

	if fields["paramval"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["paramval"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Var)
			node.Paramval = &val
		}
	}

	return
}
