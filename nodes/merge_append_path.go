// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
