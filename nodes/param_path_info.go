// Auto-generated - DO NOT EDIT

package pg_query

type ParamPathInfo struct {
	PpiReqOuter []uint32 `json:"ppi_req_outer"` /* rels supplying parameters used by path */
	PpiRows     float64  `json:"ppi_rows"`      /* estimated number of result tuples */
	PpiClauses  []Node   `json:"ppi_clauses"`   /* join clauses available from outer rels */
}
