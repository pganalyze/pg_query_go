// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * LockingClause - raw representation of FOR [NO KEY] UPDATE/[KEY] SHARE
 *		options
 *
 * Note: lockedRels == NIL means "all relations in query".  Otherwise it
 * is a list of RangeVar nodes.  (We use RangeVar mainly because it carries
 * a location field --- currently, parse analysis insists on unqualified
 * names in LockingClause.)
 */
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
