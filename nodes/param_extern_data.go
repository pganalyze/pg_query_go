// Auto-generated - DO NOT EDIT

package pg_query

type ParamExternData struct {
	Value  Datum  `json:"value"`  /* parameter value */
	Isnull bool   `json:"isnull"` /* is it NULL? */
	Pflags uint16 `json:"pflags"` /* flag bits, see above */
	Ptype  Oid    `json:"ptype"`  /* parameter's datatype, or 0 */
}
