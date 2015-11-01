// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ParamPathInfo struct {
	PpiReqOuter []uint32 `json:"ppi_req_outer"` /* rels supplying parameters used by path */
	PpiRows     float64  `json:"ppi_rows"`      /* estimated number of result tuples */
	PpiClauses  []Node   `json:"ppi_clauses"`   /* join clauses available from outer rels */
}

func (node ParamPathInfo) MarshalJSON() ([]byte, error) {
	type ParamPathInfoMarshalAlias ParamPathInfo
	return json.Marshal(map[string]interface{}{
		"PARAMPATHINFO": (*ParamPathInfoMarshalAlias)(&node),
	})
}

func (node *ParamPathInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
