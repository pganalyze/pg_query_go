// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * ParamPathInfo
 *
 * All parameterized paths for a given relation with given required outer rels
 * link to a single ParamPathInfo, which stores common information such as
 * the estimated rowcount for this parameterization.  We do this partly to
 * avoid recalculations, but mostly to ensure that the estimated rowcount
 * is in fact the same for every such path.
 *
 * Note: ppi_clauses is only used in ParamPathInfos for base relation paths;
 * in join cases it's NIL because the set of relevant clauses varies depending
 * on how the join is formed.  The relevant clauses will appear in each
 * parameterized join path's joinrestrictinfo list, instead.
 */
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
