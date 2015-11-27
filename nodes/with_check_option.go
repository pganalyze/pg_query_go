// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WithCheckOption struct {
	Viewname *string `json:"viewname"` /* name of view that specified the WCO */
	Qual     Node    `json:"qual"`     /* constraint qual to check */
	Cascaded bool    `json:"cascaded"` /* true = WITH CASCADED CHECK OPTION */
}

func (node WithCheckOption) MarshalJSON() ([]byte, error) {
	type WithCheckOptionMarshalAlias WithCheckOption
	return json.Marshal(map[string]interface{}{
		"WITHCHECKOPTION": (*WithCheckOptionMarshalAlias)(&node),
	})
}

func (node *WithCheckOption) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["viewname"] != nil {
		err = json.Unmarshal(fields["viewname"], &node.Viewname)
		if err != nil {
			return
		}
	}

	if fields["qual"] != nil {
		node.Qual, err = UnmarshalNodeJSON(fields["qual"])
		if err != nil {
			return
		}
	}

	if fields["cascaded"] != nil {
		err = json.Unmarshal(fields["cascaded"], &node.Cascaded)
		if err != nil {
			return
		}
	}

	return
}
