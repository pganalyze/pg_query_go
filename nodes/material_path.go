// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * MaterialPath represents use of a Material plan node, i.e., caching of
 * the output of its subpath.  This is used when the subpath is expensive
 * and needs to be scanned repeatedly, or when we need mark/restore ability
 * and the subpath doesn't have it.
 */
type MaterialPath struct {
	Path    Path  `json:"path"`
	Subpath *Path `json:"subpath"`
}

func (node MaterialPath) MarshalJSON() ([]byte, error) {
	type MaterialPathMarshalAlias MaterialPath
	return json.Marshal(map[string]interface{}{
		"MATERIALPATH": (*MaterialPathMarshalAlias)(&node),
	})
}

func (node *MaterialPath) UnmarshalJSON(input []byte) (err error) {
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

	if fields["subpath"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["subpath"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Subpath = &val
		}
	}

	return
}
