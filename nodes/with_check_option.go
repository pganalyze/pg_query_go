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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
