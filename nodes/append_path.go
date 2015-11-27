// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * AppendPath represents an Append plan, ie, successive execution of
 * several member plans.
 *
 * Note: it is possible for "subpaths" to contain only one, or even no,
 * elements.  These cases are optimized during create_append_plan.
 * In particular, an AppendPath with no subpaths is a "dummy" path that
 * is created to represent the case that a relation is provably empty.
 */
type AppendPath struct {
	Path     Path   `json:"path"`
	Subpaths []Node `json:"subpaths"` /* list of component Paths */
}

func (node AppendPath) MarshalJSON() ([]byte, error) {
	type AppendPathMarshalAlias AppendPath
	return json.Marshal(map[string]interface{}{
		"APPENDPATH": (*AppendPathMarshalAlias)(&node),
	})
}

func (node *AppendPath) UnmarshalJSON(input []byte) (err error) {
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

	if fields["subpaths"] != nil {
		node.Subpaths, err = UnmarshalNodeArrayJSON(fields["subpaths"])
		if err != nil {
			return
		}
	}

	return
}
