// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlannerParamItem struct {
	Item    Node `json:"item"`    /* the Var, PlaceHolderVar, or Aggref */
	ParamId int  `json:"paramId"` /* its assigned PARAM_EXEC slot number */
}

func (node PlannerParamItem) MarshalJSON() ([]byte, error) {
	type PlannerParamItemMarshalAlias PlannerParamItem
	return json.Marshal(map[string]interface{}{
		"PLANNERPARAMITEM": (*PlannerParamItemMarshalAlias)(&node),
	})
}

func (node *PlannerParamItem) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["item"] != nil {
		node.Item, err = UnmarshalNodeJSON(fields["item"])
		if err != nil {
			return
		}
	}

	if fields["paramId"] != nil {
		err = json.Unmarshal(fields["paramId"], &node.ParamId)
		if err != nil {
			return
		}
	}

	return
}
