// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["cacheId"] != nil {
		err = json.Unmarshal(fields["cacheId"], &node.CacheId)
		if err != nil {
			return
		}
	}

	if fields["hashValue"] != nil {
		err = json.Unmarshal(fields["hashValue"], &node.HashValue)
		if err != nil {
			return
		}
	}

	return
}
