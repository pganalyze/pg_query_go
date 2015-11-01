// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlanInvalItem struct {
	CacheId   int    `json:"cacheId"`   /* a syscache ID, see utils/syscache.h */
	HashValue uint32 `json:"hashValue"` /* hash value of object's cache lookup key */
}

func (node PlanInvalItem) MarshalJSON() ([]byte, error) {
	type PlanInvalItemMarshalAlias PlanInvalItem
	return json.Marshal(map[string]interface{}{
		"PLANINVALITEM": (*PlanInvalItemMarshalAlias)(&node),
	})
}

func (node *PlanInvalItem) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
