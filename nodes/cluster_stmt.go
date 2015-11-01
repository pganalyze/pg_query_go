// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ClusterStmt struct {
	Relation  *RangeVar `json:"relation"`  /* relation being indexed, or NULL if all */
	Indexname *string   `json:"indexname"` /* original index defined */
	Verbose   bool      `json:"verbose"`   /* print progress info */
}

func (node ClusterStmt) MarshalJSON() ([]byte, error) {
	type ClusterStmtMarshalAlias ClusterStmt
	return json.Marshal(map[string]interface{}{
		"CLUSTERSTMT": (*ClusterStmtMarshalAlias)(&node),
	})
}

func (node *ClusterStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
