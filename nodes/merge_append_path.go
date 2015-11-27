// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * MergeAppendPath represents a MergeAppend plan, ie, the merging of sorted
 * results from several member plans to produce similarly-sorted output.
 */
type MergeAppendPath struct {
	Path        Path    `json:"path"`
	Subpaths    []Node  `json:"subpaths"`     /* list of component Paths */
	LimitTuples float64 `json:"limit_tuples"` /* hard limit on output tuples, or -1 */
}

func (node MergeAppendPath) MarshalJSON() ([]byte, error) {
	type MergeAppendPathMarshalAlias MergeAppendPath
	return json.Marshal(map[string]interface{}{
		"MERGEAPPENDPATH": (*MergeAppendPathMarshalAlias)(&node),
	})
}

func (node *MergeAppendPath) UnmarshalJSON(input []byte) (err error) {
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

	if fields["limit_tuples"] != nil {
		err = json.Unmarshal(fields["limit_tuples"], &node.LimitTuples)
		if err != nil {
			return
		}
	}

	return
}
