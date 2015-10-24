// Auto-generated - DO NOT EDIT

package pg_query

type DefElem struct {
	Defnamespace *string       `json:"defnamespace"` /* NULL if unqualified name */
	Defname      *string       `json:"defname"`
	Arg          Node          `json:"arg"`       /* a (Value *) or a (TypeName *) */
	Defaction    DefElemAction `json:"defaction"` /* unspecified action, or SET/ADD/DROP */
}
