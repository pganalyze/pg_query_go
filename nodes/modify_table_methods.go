// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ModifyTable) MarshalJSON() ([]byte, error) {
	type ModifyTableMarshalAlias ModifyTable
	return json.Marshal(map[string]interface{}{
		"MODIFYTABLE": (*ModifyTableMarshalAlias)(&node),
	})
}

func (node *ModifyTable) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ModifyTable) Deparse() string {
	panic("Not Implemented")
}
