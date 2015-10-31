// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IndexOnlyScan) MarshalJSON() ([]byte, error) {
	type IndexOnlyScanMarshalAlias IndexOnlyScan
	return json.Marshal(map[string]interface{}{
		"INDEXONLYSCAN": (*IndexOnlyScanMarshalAlias)(&node),
	})
}

func (node *IndexOnlyScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IndexOnlyScan) Deparse() string {
	panic("Not Implemented")
}
