// Auto-generated - DO NOT EDIT

package pg_query

type FunctionScan struct {
	Scan           Scan   `json:"scan"`
	Functions      []Node `json:"functions"`      /* list of RangeTblFunction nodes */
	Funcordinality bool   `json:"funcordinality"` /* WITH ORDINALITY */
}
