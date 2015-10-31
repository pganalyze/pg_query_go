// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterTableCmd) MarshalJSON() ([]byte, error) {
	type AlterTableCmdMarshalAlias AlterTableCmd
	return json.Marshal(map[string]interface{}{
		"ALTER TABLE CMD": (*AlterTableCmdMarshalAlias)(&node),
	})
}

func (node *AlterTableCmd) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterTableCmd) Deparse() string {
	panic("Not Implemented")
}
