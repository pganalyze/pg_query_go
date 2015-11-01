// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ReplicaIdentityStmt struct {
	IdentityType byte    `json:"identity_type"`
	Name         *string `json:"name"`
}

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
