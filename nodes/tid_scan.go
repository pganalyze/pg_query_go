// Auto-generated - DO NOT EDIT

package pg_query

type TidScan struct {
	Scan     Scan   `json:"scan"`
	Tidquals []Node `json:"tidquals"` /* qual(s) involving CTID = something */
}
