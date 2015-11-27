// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["ppi_req_outer"] != nil {
		err = json.Unmarshal(fields["ppi_req_outer"], &node.PpiReqOuter)
		if err != nil {
			return
		}
	}

	if fields["ppi_rows"] != nil {
		err = json.Unmarshal(fields["ppi_rows"], &node.PpiRows)
		if err != nil {
			return
		}
	}

	if fields["ppi_clauses"] != nil {
		node.PpiClauses, err = UnmarshalNodeArrayJSON(fields["ppi_clauses"])
		if err != nil {
			return
		}
	}

	return
}
