// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type LockingClause struct {
	LockedRels []Node             `json:"lockedRels"` /* FOR [KEY] UPDATE/SHARE relations */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"` /* NOWAIT option */
}

func (node LockingClause) MarshalJSON() ([]byte, error) {
	type LockingClauseMarshalAlias LockingClause
	return json.Marshal(map[string]interface{}{
		"LOCKINGCLAUSE": (*LockingClauseMarshalAlias)(&node),
	})
}

func (node *LockingClause) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["lockedRels"] != nil {
		node.LockedRels, err = UnmarshalNodeArrayJSON(fields["lockedRels"])
		if err != nil {
			return
		}
	}

	if fields["strength"] != nil {
		err = json.Unmarshal(fields["strength"], &node.Strength)
		if err != nil {
			return
		}
	}

	if fields["noWait"] != nil {
		err = json.Unmarshal(fields["noWait"], &node.NoWait)
		if err != nil {
			return
		}
	}

	return
}
