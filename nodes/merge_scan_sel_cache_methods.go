// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MergeScanSelCache) MarshalJSON() ([]byte, error) {
	type MergeScanSelCacheMarshalAlias MergeScanSelCache
	return json.Marshal(map[string]interface{}{
		"MERGESCANSELCACHE": (*MergeScanSelCacheMarshalAlias)(&node),
	})
}

func (node *MergeScanSelCache) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MergeScanSelCache) Deparse() string {
	panic("Not Implemented")
}
