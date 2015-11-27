// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * Type "Path" is used as-is for sequential-scan paths, as well as some other
 * simple plan types that we don't need any extra information in the path for.
 * For other path types it is the first component of a larger struct.
 *
 * "pathtype" is the NodeTag of the Plan node we could build from this Path.
 * It is partially redundant with the Path's NodeTag, but allows us to use
 * the same Path type for multiple Plan types when there is no need to
 * distinguish the Plan type during path processing.
 *
 * "param_info", if not NULL, links to a ParamPathInfo that identifies outer
 * relation(s) that provide parameter values to each scan of this path.
 * That means this path can only be joined to those rels by means of nestloop
 * joins with this path on the inside.  Also note that a parameterized path
 * is responsible for testing all "movable" joinclauses involving this rel
 * and the specified outer rel(s).
 *
 * "rows" is the same as parent->rows in simple paths, but in parameterized
 * paths and UniquePaths it can be less than parent->rows, reflecting the
 * fact that we've filtered by extra join conditions or removed duplicates.
 *
 * "pathkeys" is a List of PathKey nodes (see above), describing the sort
 * ordering of the path's output rows.
 */
type Path struct {
	Parent    *RelOptInfo    `json:"parent"`     /* the relation this path can build */
	ParamInfo *ParamPathInfo `json:"param_info"` /* parameterization info, or NULL if none */

	/* estimated size/costs for path (see costsize.c for more info) */
	Rows        float64 `json:"rows"`         /* estimated number of result tuples */
	StartupCost Cost    `json:"startup_cost"` /* cost expended before fetching any tuples */
	TotalCost   Cost    `json:"total_cost"`   /* total cost (assuming all tuples fetched) */

	Pathkeys []Node `json:"pathkeys"` /* sort ordering of path's output */

	/* pathkeys is a List of PathKey nodes; see above */
}

func (node Path) MarshalJSON() ([]byte, error) {
	type PathMarshalAlias Path
	return json.Marshal(map[string]interface{}{
		"PATH": (*PathMarshalAlias)(&node),
	})
}

func (node *Path) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["parent"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["parent"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RelOptInfo)
			node.Parent = &val
		}
	}

	if fields["param_info"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["param_info"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(ParamPathInfo)
			node.ParamInfo = &val
		}
	}

	if fields["rows"] != nil {
		err = json.Unmarshal(fields["rows"], &node.Rows)
		if err != nil {
			return
		}
	}

	if fields["startup_cost"] != nil {
		err = json.Unmarshal(fields["startup_cost"], &node.StartupCost)
		if err != nil {
			return
		}
	}

	if fields["total_cost"] != nil {
		err = json.Unmarshal(fields["total_cost"], &node.TotalCost)
		if err != nil {
			return
		}
	}

	if fields["pathkeys"] != nil {
		node.Pathkeys, err = UnmarshalNodeArrayJSON(fields["pathkeys"])
		if err != nil {
			return
		}
	}

	return
}
