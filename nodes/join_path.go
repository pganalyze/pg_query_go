// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * All join-type paths share these fields.
 */
type JoinPath struct {
	Path Path `json:"path"`

	Jointype JoinType `json:"jointype"`

	Outerjoinpath *Path `json:"outerjoinpath"` /* path for the outer side of the join */
	Innerjoinpath *Path `json:"innerjoinpath"` /* path for the inner side of the join */

	Joinrestrictinfo []Node `json:"joinrestrictinfo"` /* RestrictInfos to apply to join */

	/*
	 * See the notes for RelOptInfo and ParamPathInfo to understand why
	 * joinrestrictinfo is needed in JoinPath, and can't be merged into the
	 * parent RelOptInfo.
	 */
}

func (node JoinPath) MarshalJSON() ([]byte, error) {
	type JoinPathMarshalAlias JoinPath
	return json.Marshal(map[string]interface{}{
		"JOINPATH": (*JoinPathMarshalAlias)(&node),
	})
}

func (node *JoinPath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["path"] != nil {
		err = json.Unmarshal(fields["path"], &node.Path)
		if err != nil {
			return
		}
	}

	if fields["jointype"] != nil {
		err = json.Unmarshal(fields["jointype"], &node.Jointype)
		if err != nil {
			return
		}
	}

	if fields["outerjoinpath"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["outerjoinpath"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Outerjoinpath = &val
		}
	}

	if fields["innerjoinpath"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["innerjoinpath"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Innerjoinpath = &val
		}
	}

	if fields["joinrestrictinfo"] != nil {
		node.Joinrestrictinfo, err = UnmarshalNodeArrayJSON(fields["joinrestrictinfo"])
		if err != nil {
			return
		}
	}

	return
}
