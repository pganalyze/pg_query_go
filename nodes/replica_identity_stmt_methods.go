// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ReplicaIdentityStmt) MarshalJSON() ([]byte, error) {
	type ReplicaIdentityStmtMarshalAlias ReplicaIdentityStmt
	return json.Marshal(map[string]interface{}{
		"REPLICAIDENTITYSTMT": (*ReplicaIdentityStmtMarshalAlias)(&node),
	})
}

func (node *ReplicaIdentityStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ReplicaIdentityStmt) Deparse() string {
	panic("Not Implemented")
}
