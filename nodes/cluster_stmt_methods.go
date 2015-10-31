// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ClusterStmt) Deparse() string {
	panic("Not Implemented")
}
