// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTSDictionaryStmt struct {
	Dictname []Node `json:"dictname"` /* qualified name (list of Value strings) */
	Options  []Node `json:"options"`  /* List of DefElem nodes */
}

func (node AlterTSDictionaryStmt) MarshalJSON() ([]byte, error) {
	type AlterTSDictionaryStmtMarshalAlias AlterTSDictionaryStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTSDICTIONARYSTMT": (*AlterTSDictionaryStmtMarshalAlias)(&node),
	})
}

func (node *AlterTSDictionaryStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
