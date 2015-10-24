// Auto-generated - DO NOT EDIT

package pg_query

type QualCost struct {
	Startup  Cost `json:"startup"`   /* one-time cost */
	PerTuple Cost `json:"per_tuple"` /* per-evaluation cost */
}
